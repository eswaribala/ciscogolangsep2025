package store

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/vault/api"
	"github.com/segmentio/kafka-go"
)

type Site struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name" gorm:"type:varchar(100);uniqueIndex;not null"`
	Location string `json:"location" gorm:"type:varchar(100);not null"`
	Status   bool   `json:"status" gorm:"not null;default:true"`
}

const (
	kafkaTopic   = "site-events"
	kafkaBroker  = "kafka:9092" // or host.docker.internal:9092 if calling from another container
	publishEvent = "site_fetched"
)

// Reusable writer (initialize once in init() or main())
var siteWriter *kafka.Writer

func EnsureTopic() error {
	println("Ensuring kafka topic")
	// 1) Dial any broker
	b, err := kafka.Dial("tcp", kafkaBroker)
	if err != nil {
		return fmt.Errorf("dial broker %s: %w", kafkaBroker, err)
	}
	defer b.Close()

	// 2) Ask who the controller is
	ctrlInfo, err := b.Controller()
	if err != nil {
		return fmt.Errorf("get controller: %w", err)
	}
	ctrlAddr := net.JoinHostPort(ctrlInfo.Host, strconv.Itoa(ctrlInfo.Port))

	// 3) Dial controller and set a deadline (replacement for context)
	ctrl, err := kafka.Dial("tcp", ctrlAddr)
	if err != nil {
		return fmt.Errorf("dial controller %s: %w", ctrlAddr, err)
	}
	defer ctrl.Close()
	_ = ctrl.SetDeadline(time.Now().Add(10 * time.Second))

	// 4) Create topic (idempotent: ignore "already exists")
	err = ctrl.CreateTopics(kafka.TopicConfig{
		Topic:             kafkaTopic,
		NumPartitions:     1,
		ReplicationFactor: 1,
	})
	if err != nil {
		if strings.Contains(strings.ToLower(err.Error()), "already exist") {
			log.Printf("[kafka] topic %q already exists", kafkaTopic)
			return nil
		}
		return fmt.Errorf("create topic %q: %w", kafkaTopic, err)
	}
	log.Printf("[kafka] created topic %q (partitions=%d, rf=%d)", kafkaTopic, 1, 1)
	return nil
}

func InitKafkaWriters() {
	siteWriter = &kafka.Writer{
		Addr:                   kafka.TCP(kafkaBroker),
		Topic:                  kafkaTopic,
		Balancer:               &kafka.Hash{},
		RequiredAcks:           kafka.RequireAll,
		BatchTimeout:           100 * time.Millisecond,
		AllowAutoTopicCreation: false, // keep explicit creation for stability
		Logger:                 log.New(os.Stdout, "[kafka] ", log.LstdFlags),
		ErrorLogger:            log.New(os.Stderr, "[kafka-err] ", log.LstdFlags),
	}
}

// Call this on app shutdown
func CloseKafkaWriters() {
	if siteWriter != nil {
		_ = siteWriter.Close()
	}
}

// CreateSite godoc
// @Summary      Create a new site
// @Description  Adds a new site
// @Tags         sites
// @Accept       json
// @Produce      json
// @Param        site  body      Site  true  "Site to create"
// @Success      201    {object}  Site
// @Failure      400    {object}  map[string]string "Invalid input"
// @Router       /sites/v1.0 [post]
func SaveSiteInfo(writer http.ResponseWriter, request *http.Request) {
	//store the site info to db
	db := MySQLConnectionHelper()
	//defer db.Close()

	//parse the request body
	var site Site
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	//save the site info to db
	result := db.Create(&site)
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(site)

}

// GetAllSites godoc
// @Summary      Get all sites
// @Description  Returns list of sites
// @Tags         sites
// @Accept       json
// @Produce      json
// @Success      200  {array}   Site
// @Router       /sites/v1.0 [get]
func GetAllSiteInfo(writer http.ResponseWriter, request *http.Request) {
	var sites []Site
	db := MySQLConnectionHelper()
	result := db.Find(&sites)
	if result.Error != nil {

		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(sites)

}

// GetSiteById godoc
// @Summary Get details of requested site
// @Description Get details of requestesd site
// @Tags sites
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Site"
// @Success 200 {object} Site
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Site not found"
// @Router /sites/v1.0/{id} [get]
func GetSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	result := db.First(&site, request.URL.Query().Get("id"))
	if result.Error != nil {

		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	//return the site info to client
	json.NewEncoder(writer).Encode(site)

}

// UpdateSite godoc
// @Summary Update existing site
// @Description Update existing site with the input payload
// @Tags sites
// @Accept  json
// @Produce  json
// @Param site body Site true "Update site"
// @Success 200 {object} Site
// @Router /sites/v1.0 [put]
func UpdateSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	err := json.NewDecoder(request.Body).Decode(&site)
	if err != nil {
		http.Error(writer, err.Error(), http.StatusBadRequest)
		return
	}
	result := db.Model(&site).Updates(Site{Name: site.Name, Location: site.Location, Status: site.Status})
	//send result
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode(site)

}

// DeleteSiteById godoc
// @Summary Delete requested site
// @Description Delete requested site
// @Tags sites
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Site"
// @Success 200 {object} Site
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Site not found"
// @Router /sites/v1.0/{id} [delete]
func DeleteSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	db := MySQLConnectionHelper()
	// Extract from path instead of query
	idStr := request.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil || id <= 0 {
		http.Error(writer, `{"error":"Invalid ID supplied"}`, http.StatusBadRequest)
		return
	}
	println("ID to delete:", id)
	result := db.Where("id=?", id).Delete(&Site{})
	if result.Error != nil {
		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	json.NewEncoder(writer).Encode("Site deleted successfully")

}

// GetSiteById godoc
// @Summary Get details of requested site
// @Description Get details of requestesd site
// @Tags sites
// @Accept  json
// @Produce  json
// @Param id path int true "ID of the Site"
// @Success 200 {object} Site
// @Failure 400 {object} map[string]string "Invalid ID supplied"
// @Failure 404 {object} map[string]string "Site not found"
// @Router /sites/v1.0/kafka/{id} [get]
func PublishSiteInfoByID(writer http.ResponseWriter, request *http.Request) {
	var site Site
	db := MySQLConnectionHelper()
	result := db.First(&site, request.URL.Query().Get("id"))
	if result.Error != nil {

		http.Error(writer, result.Error.Error(), http.StatusInternalServerError)
		return
	}
	// 3) Publish to Kafka (best-effort, with timeout)
	go func(s Site) {
		// encode payload
		body, err := json.Marshal(struct {
			Event string `json:"event"`
			Site  Site   `json:"site"`
			TS    string `json:"ts"`
		}{
			Event: publishEvent,
			Site:  s,
			TS:    time.Now().Format(time.RFC3339Nano),
		})
		if err != nil {
			log.Printf("[kafka] marshal error: %v", err)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
		defer cancel()

		if err := siteWriter.WriteMessages(ctx, kafka.Message{
			Key:   []byte("site-" + strconv.Itoa(int(s.ID))), // stable partitioning by site id
			Value: body,
			Headers: []kafka.Header{
				{Key: "source", Value: []byte("sites-api")},
				{Key: "op", Value: []byte("read")},
			},
		}); err != nil {
			log.Printf("[kafka] publish error: %v", err)
		}
	}(site)

	// 4) Respond to client
	writer.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(writer).Encode(site)

}

func VaultConnection() (string, string) {
	println("Connecting to vault")
	token := "s.zUxDVRPsUVW3mX9A76gwqpZH"
	vaultAddr := "http://vault:8200"

	client, err := api.NewClient(&api.Config{Address: vaultAddr, HttpClient: &http.Client{Timeout: 100 * time.Second}})
	if err != nil {
		println("Error connecting to vault:", err)
	}
	client.SetToken(token)
	println("Connected to vault")
	// If your secrets engine at mount "secret/" is KV v2:
	kv := client.KVv2("secret")

	// name is "mysqlsecret" (the entry you see in UI)
	secret, err := kv.Get(context.Background(), "mysqlsecret")
	if err != nil {
		log.Fatalf("vault read: %v", err)
	}
	if secret == nil || secret.Data == nil {
		log.Fatal("vault: empty secret or no data")
	}

	// Keys: mysqlusername, mysqlpassword (from your screenshot)
	user, _ := secret.Data["mysqlusername"].(string)
	pass, _ := secret.Data["mysqlpassword"].(string)

	fmt.Println("username:", user)
	fmt.Println("password:", pass) // ⚠️ don’t log secrets in real apps
	return user, pass
}

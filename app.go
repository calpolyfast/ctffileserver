package main

import (
	"log"
	"time"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/file/{name}", file)

	log.Fatal(http.ListenAndServe(":8080", router))
}

// Download a file based on URL parameters
func file(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		vars := mux.Vars(r)
		fileName := vars["name"]
		bucket := r.URL.Query().Get("bucket")

		// Check if bucket name given
		if len(bucket) <= 0 {
			bucket = "calpolyfast-ctf-fall16"
		}

		// Create new s3 instance to work with
		svc := s3.New(session.New(&aws.Config{Region: aws.String("us-west-2")}))

		// Check if object exists
		_, err := svc.HeadObject(&s3.HeadObjectInput{
			Bucket: aws.String(bucket),
		    Key:    aws.String(fileName),
		})
		if err != nil {
			http.Error(w, "File Does Not Exist", http.StatusBadRequest)
			return
		}

		// Create object request to perform actions on
		req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		    Bucket: aws.String(bucket),
		    Key:    aws.String(fileName),
		})

		// Get url for file that expires after 30 minutes
		urlStr, err := req.Presign(30 * time.Minute)
		if err != nil {
		    log.Println("Failed to sign request", err)
		}

		http.Redirect(w, r, urlStr, http.StatusSeeOther)
	}
}

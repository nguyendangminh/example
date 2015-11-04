package main 

import "fmt"
import "log"
import "os"
import "io"
import "github.com/mulander/gojenkins"


func main() {
	fmt.Println("Running")

	// configuration
	baseurl := "http://some.jenkins.build.server.example.com"
	username := "username"
	token := "secret-api-token"
	testjob := "TestJob"
	buildString := "lastSuccessfulBuild"

	jenkins := &gojenkins.Jenkins{
	  Baseurl: baseurl,
	}
	jenkins.SetAuth(username, token)

	// List all Jenkins jobs
	jobs, err := jenkins.Jobs()
	if err != nil {
	    log.Fatal(err)
	}

	// List all artifacts of a job
	artifacts, err := jenkins.Artifacts(jobs[testjob], buildString)
	if err != nil {
	    log.Fatal(err)
	}

	// For every job artifact
	for _, artifact := range artifacts {
	    out, err := os.Create(artifact.FileName)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer out.Close()

	    // perform a download receiving a reader
	    a, err := jenkins.Download(jobs[testjob], buildString, artifact)
	    if err != nil {
	        log.Fatal(err)
	    }
	    defer a.Close()

	    // store it to disk using io/ioutil
	    _, err = io.Copy(out, a)
	    if err != nil {
	        log.Fatal(err)
	    }
	}
}
package node

import "net/http"

func HealthCheck(w http.ResponseWriter, r *http.Request) {}
func AddNode(w http.ResponseWriter, r *http.Request)     {}
func RemoveNode(w http.ResponseWriter, r *http.Request)  {}
func GetNode(w http.ResponseWriter, r *http.Request)     {}
func GetNodes(w http.ResponseWriter, r *http.Request)    {}

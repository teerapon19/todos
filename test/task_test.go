package test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/teerapon19/todos/model"
)

type TaskResponse struct {
	Status  string
	Message string
}

type GetAllTaskRespone struct {
	TaskResponse
	Tasks []model.Task
}

type GetSingleTaskRespone struct {
	TaskResponse
	Task model.Task
}

func Test_CreateNewTask(t *testing.T) {
	task := []byte(`{"title":"Test my task", "description": "This's test."}`)
	resp, err := http.Post("http://localhost:8080/task", "application/json", bytes.NewBuffer(task))
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	response := TaskResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response.Status)
	assert.Equal(t, 200, resp.StatusCode)
}

func Test_GetAllTaskAndGetSingle(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/task/all")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getAllTaskRespone := GetAllTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getAllTaskRespone)

	assert.GreaterOrEqual(t, len(getAllTaskRespone.Tasks), 1, "Get all task length")
	assert.Equal(t, 200, resp.StatusCode, "Get all task status code")

	lastTaskID := getAllTaskRespone.Tasks[len(getAllTaskRespone.Tasks)-1].ID.Hex()
	url := fmt.Sprintf("http://localhost:8080/task/%s", lastTaskID)

	resp, err = http.Get(url)
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getSingleTaskRespone := GetSingleTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getSingleTaskRespone)

	assert.Equal(t, "success", getSingleTaskRespone.Status, "Get single task success")
	assert.Equal(t, 200, resp.StatusCode, "Get single status code")
}

func Test_GetAllTaskAndMarkOne(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/task/all")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getAllTaskRespone := GetAllTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getAllTaskRespone)

	assert.GreaterOrEqual(t, len(getAllTaskRespone.Tasks), 1, "Get all task length")
	assert.Equal(t, 200, resp.StatusCode, "Get all task status code")

	lastTaskID := getAllTaskRespone.Tasks[len(getAllTaskRespone.Tasks)-1].ID.Hex()
	url := fmt.Sprintf("http://localhost:8080/task/mark/%s", lastTaskID)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	req.Header.Add("Content-type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}

	response := TaskResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response.Status, "Mark task success")
	assert.Equal(t, 200, resp.StatusCode, "Mark status code")
}

func Test_GetAllTaskAndUnmarkOne(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/task/all")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getAllTaskRespone := GetAllTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getAllTaskRespone)

	assert.GreaterOrEqual(t, len(getAllTaskRespone.Tasks), 1, "Get all task length")
	assert.Equal(t, 200, resp.StatusCode, "Get all task status code")

	lastTaskID := getAllTaskRespone.Tasks[len(getAllTaskRespone.Tasks)-1].ID.Hex()
	url := fmt.Sprintf("http://localhost:8080/task/unmark/%s", lastTaskID)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", url, nil)
	req.Header.Add("Content-type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}

	response := TaskResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response.Status, "Unmark task success")
	assert.Equal(t, 200, resp.StatusCode, "Unmark status code")
}

func Test_GetAllTaskAndEditOne(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/task/all")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getAllTaskRespone := GetAllTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getAllTaskRespone)

	assert.GreaterOrEqual(t, len(getAllTaskRespone.Tasks), 1, "Get all task length")
	assert.Equal(t, 200, resp.StatusCode, "Get all task status code")

	msgJson := fmt.Sprintf(`{"id": "%s", "title":"Test my task edited", "description": "This's test."}`, getAllTaskRespone.Tasks[len(getAllTaskRespone.Tasks)-1].ID.Hex())
	task := []byte(msgJson)

	client := &http.Client{}
	req, err := http.NewRequest("PUT", "http://localhost:8080/task", bytes.NewBuffer(task))
	req.Header.Add("Content-type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}

	response := TaskResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response.Status, "Edit task success")
	assert.Equal(t, 200, resp.StatusCode, "Edit status code")
}

func Test_GetAllTaskAndDeleteOne(t *testing.T) {
	resp, err := http.Get("http://localhost:8080/task/all")
	if err != nil {
		t.Error(err)
	}
	defer resp.Body.Close()

	getAllTaskRespone := GetAllTaskRespone{}
	json.NewDecoder(resp.Body).Decode(&getAllTaskRespone)

	assert.GreaterOrEqual(t, len(getAllTaskRespone.Tasks), 1, "Get all task length")
	assert.Equal(t, 200, resp.StatusCode, "Get all task status code")

	lastTaskID := getAllTaskRespone.Tasks[len(getAllTaskRespone.Tasks)-1].ID.Hex()
	url := fmt.Sprintf("http://localhost:8080/task/%s", lastTaskID)

	client := &http.Client{}
	req, err := http.NewRequest("DELETE", url, nil)
	req.Header.Add("Content-type", "application/json")
	resp, err = client.Do(req)
	if err != nil {
		t.Error(err)
	}

	response := TaskResponse{}
	json.NewDecoder(resp.Body).Decode(&response)

	assert.Equal(t, "success", response.Status, "Delete task success")
	assert.Equal(t, 200, resp.StatusCode, "Delete status code")
}

package user

import (
    "bytes"
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/gin-gonic/gin"
    . "github.com/smartystreets/goconvey/convey"
)

func TestLogin(t *testing.T) {
    // Set Gin to Test Mode
    gin.SetMode(gin.TestMode)

    Convey("Given a Gin router with the Login route", t, func() {
        router := gin.Default()
        router.POST("/login", Login)

        Convey("When a valid login request is made", func() {
            body := `{"username": "testuser", "password": "testpass"}`
            req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(body)))
            So(err, ShouldBeNil)
            req.Header.Set("Content-Type", "application/json")

            // Record the HTTP response
            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)

            Convey("Then the response should be 200 OK", func() {
                So(w.Code, ShouldEqual, http.StatusOK)
            })

            Convey("And the response body should indicate successful login", func() {
                expectedBody := `{"message":"Logged In"}`
                So(w.Body.String(), ShouldEqual, expectedBody)
            })
        })

        Convey("When an invalid login request is made", func() {
            body := `{"username": "testuser"}`
            req, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer([]byte(body)))
            So(err, ShouldBeNil)
            req.Header.Set("Content-Type", "application/json")

            // Record the HTTP response
            w := httptest.NewRecorder()
            router.ServeHTTP(w, req)

            Convey("Then the response should be 400 Bad Request", func() {
                So(w.Code, ShouldEqual, http.StatusBadRequest)
            })

            Convey("And the response body should indicate an error", func() {
                expectedBody := `{"error":"Invalid input"}`
                So(w.Body.String(), ShouldEqual, expectedBody)
            })
        })
    })
}
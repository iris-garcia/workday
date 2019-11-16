// Unit tests
package api_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/DATA-DOG/go-sqlmock"

	. "github.com/iris-garcia/workday/api"
)

// API Router
var _ = Describe("API Router", func() {
	Describe("When requesting the endpoint GET /", func() {
		db, _, _ := sqlmock.New()
		router := GinRouter(db)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		It("Should return a 200", func() {
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Should return the expected JSON", func() {
			Expect(w.Body.String()).To(Equal(`{"status":"OK"}`))
		})
	})

	Describe("When requesting the endpoint GET /status", func() {
		db, _, _ := sqlmock.New()
		router := GinRouter(db)
		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/status", nil)
		router.ServeHTTP(w, req)

		It("Should return a 200", func() {
			Expect(w.Code).To(Equal(http.StatusOK))
		})

		It("Should return the expected JSON", func() {
			Expect(w.Body.String()).To(Equal(`{"status":"OK"}`))
		})
	})

	Describe("When requesting the endpoint GET /employees", func() {
		req, _ := http.NewRequest("GET", "/employees", nil)
		Context("And reading employees from DB is successful", func() {
			db, mock, _ := sqlmock.New()
			rows := mock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Password"})
			mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)
			router := GinRouter(db)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			It("Should return a 200 code", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
		Context("And reading employees from DB is not successful", func() {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnError(errors.New("asd"))
			router := GinRouter(db)
			w := httptest.NewRecorder()
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
		})
	})

	Describe("When requesting the endpoint GET /employees:id", func() {
		Context("And the id param is not valid", func() {
			db, _, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/employees/a", nil)
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
			It("Should return an error", func() {
				Expect(w.Body.String()).To(Equal(`{"error":"strconv.Atoi: parsing \"a\": invalid syntax"}`))
			})
		})
		Context("And the employee id does not exists", func() {
			db, mock, _ := sqlmock.New()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"})
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnRows(rows)

			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/employees/2", nil)
			router.ServeHTTP(w, req)
			It("Should return a 404 code", func() {
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
			It("Should return an empty json", func() {
				Expect(w.Body.String()).To(Equal(`{}`))
			})
		})

		Context("And the DB returns an error", func() {
			db, mock, _ := sqlmock.New()
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnError(errors.New("DB error"))

			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/employees/2", nil)
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
		})

		Context("And the employee id exists", func() {
			db, mock, _ := sqlmock.New()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
				AddRow(1, "Iris", "Garcia", 1, "secret")
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnRows(rows)
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/employees/2", nil)
			router.ServeHTTP(w, req)

			It("Should return a 200 code", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
			})
			It("Should return a JSON with the employee", func() {
				expected := Employee{ID: 1, Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}
				var body Employee
				json.NewDecoder(w.Body).Decode(&body)
				Expect(body).To(Equal(expected))
			})
		})
	})

	Describe("When requesting the endpoint DELETE /employees:id", func() {
		Context("And the id param is not valid", func() {
			db, _, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/employees/a", nil)
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
			It("Should return an error", func() {
				Expect(w.Body.String()).To(Equal(`{"error":"strconv.Atoi: parsing \"a\": invalid syntax"}`))
			})
		})
		Context("And the employee is not deleted in the DB", func() {
			db, mock, _ := sqlmock.New()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
				AddRow(1, "Iris", "Garcia", 1, "secret")
			mock.ExpectQuery("^DELETE FROM employee WHERE id=(.+)").WillReturnRows(rows)
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/employees/2", nil)
			router.ServeHTTP(w, req)

			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
		})
		Context("And the employee is deleted", func() {
			db, mock, _ := sqlmock.New()
			mock.ExpectExec("DELETE FROM employee").WithArgs(1).
				WillReturnResult(sqlmock.NewResult(1, 1))
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("DELETE", "/employees/1", nil)
			router.ServeHTTP(w, req)
			It("Should return a 200 code", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
			})
			It("Should return a JSON with message OK", func() {
				Expect(w.Body.String()).To(Equal(`{"message":"OK"}`))
			})
		})
	})

	Describe("When requesting the endpoint PUT /employees:id", func() {
		Context("And the id param is not valid", func() {
			db, _, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("PUT", "/employees/a", nil)
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
			It("Should return an error", func() {
				Expect(w.Body.String()).To(Equal(`{"error":"strconv.Atoi: parsing \"a\": invalid syntax"}`))
			})
		})
		Context("And the post data is not valid", func() {
			db, _, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garcia"}`)
			req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)
			It("Should return a 204 code", func() {
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
		Context("And the employee id does not exists", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnError(errors.New("DB error"))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garciaa", "role": 1, "password": "new"}`)
			req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)
			It("Should return a 404 code", func() {
				Expect(w.Code).To(Equal(http.StatusNotFound))
			})
		})
		Context("And there is a DB transaction error", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
				AddRow(1, "Iris", "Garcia", 1, "secret")
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnRows(rows)
			mock.ExpectExec("^UPDATE employee SET (.+) WHERE id=(.+)").WillReturnError(errors.New("DB error"))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garciaa", "role": 1, "password": "new"}`)
			req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
		})
		Context("And the employee id exists more than one time", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
				AddRow(1, "Iris", "Garcia", 1, "secret")
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnRows(rows)
			mock.ExpectExec("^UPDATE employee SET (.+) WHERE id=(.+)").
				WillReturnResult(sqlmock.NewResult(0, 2))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garcia", "role": 1, "password": "new"}`)
			req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			It("Should edit just 1 row", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
				Expect(w.Body.String()).To(Equal(`{"error":"2 rows updated"}`))
			})
			It("Should return a 200 code", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
		Context("And the employee id exists just once", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
				AddRow(1, "Iris", "Garcia", 1, "secret")
			mock.ExpectQuery("^SELECT (.+) FROM employee WHERE id=(.+)").WillReturnRows(rows)
			mock.ExpectExec("^UPDATE employee SET (.+) WHERE id=(.+)").
				WillReturnResult(sqlmock.NewResult(0, 1))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garciaa", "role": 1, "password": "new"}`)
			req, _ := http.NewRequest("PUT", "/employees/1", bytes.NewBuffer(jsonStr))
			router.ServeHTTP(w, req)

			It("Should return a 200 code", func() {
				Expect(w.Code).To(Equal(http.StatusOK))
			})
		})
	})

	Describe("When requesting the endpoint POST /employees", func() {
		Context("And the post data is not valid", func() {
			db, _, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garcia"}`)
			req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)

			It("Should return a 204 code", func() {
				Expect(w.Code).To(Equal(http.StatusNoContent))
			})
		})
		Context("And inserting the employee in the DB is not successful", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			mock.ExpectQuery("^INSERT INTO employee (firstname, lastname, role, password) values(.+)").WillReturnError(errors.New("DB error"))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garcia", "role": 1, "password": "12345"}`)
			req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			It("Should return a 500 code", func() {
				Expect(w.Code).To(Equal(http.StatusInternalServerError))
			})
			It("Should return an error message", func() {
				Expect(w.Body.String()).To(MatchRegexp("DB error"))
			})
		})
		Context("And inserting the employee in the DB is successful", func() {
			db, mock, _ := sqlmock.New()
			router := GinRouter(db)
			w := httptest.NewRecorder()
			mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
				WillReturnResult(sqlmock.NewResult(1, 1))
			var jsonStr = []byte(`{"firstname":"Iris", "lastname": "Garcia", "role": 1, "password": "secret"}`)
			req, _ := http.NewRequest("POST", "/employees", bytes.NewBuffer(jsonStr))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			It("Should return a 201 code", func() {
				Expect(w.Code).To(Equal(http.StatusCreated))
			})
			It("Should return the employee's ID", func() {
				Expect(w.Body.String()).To(Equal(`{"id":1}`))
			})
		})
	})
})

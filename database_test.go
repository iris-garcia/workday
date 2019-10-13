package workday_test

import (
	"database/sql"

	"github.com/DATA-DOG/go-sqlmock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/iris-garcia/workday"
)

type MyCustomError struct {
	message string
}

func (e MyCustomError) Error() string {
	return e.message
}

var _ = Describe("Database", func() {
	It("Should return a sql.DB struct when calling ConnectDB", func() {
		cfg := DBConfig{}
		db, err := ConnectDB(cfg)

		Ω(err).ShouldNot(HaveOccurred())
		Ω(db).Should(BeAssignableToTypeOf(&sql.DB{}))
	})

	It("Should create an Employee", func() {
		db, mock, err := sqlmock.New()
		defer db.Close()

		Ω(err).ShouldNot(HaveOccurred())
		mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
			WillReturnResult(sqlmock.NewResult(1, 1))

		emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}

		id, rows, err := CreateEmployee(db, emp)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(id).Should(Equal(uint(1)))
		Ω(rows).Should(Equal(uint(1)))
		// make sure that all expectations were met
		err = mock.ExpectationsWereMet()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("Should fail when the database transaction fails", func() {
		db, mock, err := sqlmock.New()
		defer db.Close()

		Ω(err).ShouldNot(HaveOccurred())
		mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
			WillReturnError(MyCustomError{message: "Error running db.Exec"})

		emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}
		_, _, err = CreateEmployee(db, emp)

		Ω(err).Should(HaveOccurred())

		// make sure that all expectations were met
		err = mock.ExpectationsWereMet()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("Should fail when the insert of new employee does not return a lastID equal to 1", func() {
		db, mock, err := sqlmock.New()
		defer db.Close()

		Ω(err).ShouldNot(HaveOccurred())
		mock.ExpectExec("INSERT INTO employee").WithArgs("Iris", "Garcia", 1, "secret").
			WillReturnResult(sqlmock.NewErrorResult(MyCustomError{message: "Error inserting"}))

		emp := Employee{Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"}
		id, rows, err := CreateEmployee(db, emp)

		Ω(err).Should(HaveOccurred())
		Ω(id).ShouldNot(Equal(uint(1)))
		Ω(rows).ShouldNot(Equal(uint(1)))

		// make sure that all expectations were met
		err = mock.ExpectationsWereMet()
		Ω(err).ShouldNot(HaveOccurred())
	})

	It("Should retrieve all the employees from the DB", func() {
		db, mock, err := sqlmock.New()
		defer db.Close()

		Ω(err).ShouldNot(HaveOccurred())

		// expect required DB actions
		rows := sqlmock.NewRows([]string{"ID", "Firstname", "Lastname", "Role", "Passwordd"}).
			AddRow(1, "Iris", "Garcia", 1, "secret").
			AddRow(2, "Name", "Lastname", 2, "changeme")

		mock.ExpectQuery("^SELECT (.+) FROM employee$").WillReturnRows(rows)

		expected := []Employee{
			Employee{ID: 1, Firstname: "Iris", Lastname: "Garcia", Role: 1, Password: "secret"},
			Employee{ID: 2, Firstname: "Name", Lastname: "Lastname", Role: 2, Password: "changeme"},
		}

		emps, err := GetAllEmployees(db)

		Ω(emps).Should(Equal(expected))
	})
})

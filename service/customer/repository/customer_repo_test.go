package repository

import (
	"testing"

	"itmx_test/domain"
	"itmx_test/service/entity"
	"itmx_test/util"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestCreate(t *testing.T) {
	// Mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	// Expectation for the sqlite version check
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	repo := NewCustomerRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Create(&entity.Customer{ID: util.GenerateUuid(), Name: "test", Age: 11})
		assert.NoError(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("INSERT INTO").WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.Create(&entity.Customer{ID: util.GenerateUuid(), Name: "test", Age: 11})
		assert.Error(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestFindByID(t *testing.T) {
	// Mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	// Expectation for the sqlite version check
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	repo := NewCustomerRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {
		// Setup expectations
		mock.ExpectQuery("SELECT").WithArgs("test-id").WillReturnRows(sqlmock.NewRows([]string{"id", "created_at", "updated_at", "name", "age"}).AddRow("test-id", "2024-04-25 22:17:32.000", "2024-04-25 22:17:32.000", "test_name", 30))

		customer, err := repo.FindByID("test-id")
		assert.NoError(t, err)
		assert.NotNil(t, customer)
		assert.Equal(t, "test-id", customer.ID)
		assert.Equal(t, "test_name", customer.Name)
		assert.Equal(t, 30, customer.Age)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {
		// Setup expectations
		mock.ExpectQuery("SELECT").WithArgs("non-existent-id").WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.FindByID("non-existent-id")
		assert.Equal(t, domain.ErrNotFound, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestUpdate(t *testing.T) {
	// Mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	// Expectation for the sqlite version check
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	repo := NewCustomerRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.Update(&entity.Customer{ID: util.GenerateUuid(), Name: "test", Age: 11})
		assert.NoError(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE").WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.Update(&entity.Customer{ID: util.GenerateUuid(), Name: "test", Age: 11})
		assert.Error(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

func TestDeleteByID(t *testing.T) {
	// Mock database
	sqlDB, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("An error '%s' was not expected when opening a stub database connection", err)
	}
	defer sqlDB.Close()

	// Expectation for the sqlite version check
	mock.ExpectQuery("select sqlite_version()").WillReturnRows(sqlmock.NewRows([]string{"version"}).AddRow("3.31.1"))

	dialector := sqlite.Dialector{Conn: sqlDB}
	gormDB, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		t.Fatalf("Failed to open gorm database: %v", err)
	}

	repo := NewCustomerRepository(gormDB)

	// Success case
	t.Run("success", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `customers` SET `deleted_at`.*").WithArgs(sqlmock.AnyArg(), "test-id").WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.DeleteByID("test-id")
		assert.NoError(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})

	// Failure case
	t.Run("failure", func(t *testing.T) {
		// Setup expectations
		mock.ExpectBegin()
		mock.ExpectExec("UPDATE `customers` SET `deleted_at`.*").WithArgs(sqlmock.AnyArg(), "test-id").WillReturnError(gorm.ErrInvalidData)
		mock.ExpectRollback()

		err := repo.DeleteByID("test-id")
		assert.Error(t, err)

		// Ensure all expectations were met
		assert.NoError(t, mock.ExpectationsWereMet())
	})
}

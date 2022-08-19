package repository

import (
	"calendar/internals/models"
	"calendar/internals/validate"
	"testing"
	"time"

	"github.com/jackc/pgx/v4"

	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestPostgresUserRepository_CreateUser(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	userUUID := uuid.NewString()
	login := "myLogin"
	passwordHash := "passHash"
	createdAt := time.Now()

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	expected := pgxpoolmock.NewRow(userUUID, login, passwordHash, createdAt, false)
	mockPool.EXPECT().
		QueryRow(gomock.Any(), gomock.Any(), gomock.Eq(login), gomock.Eq(passwordHash)).
		Return(expected)

	repo := NewPostgresUserRepository(mockPool)
	user := models.User{
		Login:        login,
		PasswordHash: passwordHash,
	}

	created := models.User{
		Id:           userUUID,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	res, err := repo.CreateUser(user)

	// then
	assert.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, created, res)
}

func TestPostgresUserRepository_GetUserByLogin1(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	userUUID := uuid.NewString()
	login := "me"
	passwordHash := "hashHere"
	createdAt := time.Now()

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "login", "password_hash", "created_at", "is_deleted"}
	pgxRows := pgxpoolmock.NewRows(columns).AddRow(userUUID, login, passwordHash, createdAt, false).ToPgxRows()

	mockPool.EXPECT().
		Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).
		Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)

	expected := models.User{
		Id:           userUUID,
		Login:        login,
		PasswordHash: passwordHash,
		CreatedAt:    createdAt,
		IsDeleted:    false,
	}

	res, err := repo.GetUserByLogin(login)

	// then
	assert.NotNil(t, res)
	assert.NoError(t, err)
	assert.Equal(t, expected, res)
}

func TestPostgresUserRepository_GetUserByLogin2(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	//userUUID := uuid.NewString()
	login := "me"
	//passwordHash := "hashHere"
	//createdAt := time.Now()
	//
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	//columns := []string{"id", "login", "password_hash", "created_at", "is_deleted"}
	//pgxRows := pgxpoolmock.NewRows(columns).AddRow(userUUID, login, passwordHash, createdAt, false).ToPgxRows()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).Return(nil, pgx.ErrNoRows)

	repo := NewPostgresUserRepository(mockPool)

	//expected := models.User{
	//	Id:           userUUID,
	//	Login:        login,
	//	PasswordHash: passwordHash,
	//	CreatedAt:    createdAt,
	//	IsDeleted:    false,
	//}

	_, err := repo.GetUserByLogin(login)
	assert.Error(t, err)
	// then
	//assert.NotNil(t, res)
	//assert.NoError(t, err)
	//assert.Equal(t, expected, res)
}

func TestPostgresUserRepository_GetUserByLogin3(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	//userUUID := uuid.NewString()
	login := "me"
	//passwordHash := "hashHere"
	//createdAt := time.Now()
	//
	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	columns := []string{"id", "login", "password_hash", "created_at", "is_deleted"}
	pgxRows := pgxpoolmock.NewRows(columns).ToPgxRows()

	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(login)).Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)

	//expected := models.User{
	//	Id:           userUUID,
	//	Login:        login,
	//	PasswordHash: passwordHash,
	//	CreatedAt:    createdAt,
	//	IsDeleted:    false,
	//}

	_, err := repo.GetUserByLogin(login)
	assert.Error(t, err)
	// then
	//assert.NotNil(t, res)
	//assert.NoError(t, err)
	//assert.Equal(t, expected, res)
}

//func TestPostgresUserRepository_GetUserByLogin(t *testing.T) {
//	type fields struct {
//		client database.Client
//	}
//	type args struct {
//		login string
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    models.User
//		wantErr assert.ErrorAssertionFunc
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PostgresUserRepository{
//				client: tt.fields.client,
//			}
//			got, err := p.GetUserByLogin(tt.args.login)
//			if !tt.wantErr(t, err, fmt.Sprintf("GetUserByLogin(%v)", tt.args.login)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "GetUserByLogin(%v)", tt.args.login)
//		})
//	}
//}

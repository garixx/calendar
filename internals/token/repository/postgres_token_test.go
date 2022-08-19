package repository

import (
	"calendar/internals/jwt"
	"calendar/internals/models"
	"calendar/internals/validate"
	"testing"

	"github.com/jackc/pgx/v4"

	"github.com/chrisyxlee/pgxpoolmock"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestPostgresTokenRepository_CreateToken(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	expected := pgxpoolmock.NewRow(jwtToken)
	mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(expected)

	repo := NewPostgresUserRepository(mockPool)

	res, err := repo.CreateToken(token)

	// then
	assert.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, token, res)
}

func TestPostgresTokenRepository_GetToken(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	pgxRows := pgxpoolmock.NewRows([]string{"token"}).AddRow(jwtToken).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)
	res, err := repo.GetToken(token)

	// then
	assert.NoError(t, err)
	assert.NoError(t, validate.Struct(res))
	assert.Equal(t, token, res)
}

func TestPostgresTokenRepository_GetToken2(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(nil, pgx.ErrNoRows)

	repo := NewPostgresUserRepository(mockPool)
	_, err = repo.GetToken(token)

	// then
	assert.Error(t, err)
	//assert.NoError(t, validate.Struct(res))
	//assert.Equal(t, token, res)
}

func TestPostgresTokenRepository_GetToken3(t *testing.T) {
	t.Parallel()
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// given
	jwtToken, err := jwt.GenerateToken("me")
	require.NoError(t, err, "jwt token should be generated")

	token := models.Token{
		Token: jwtToken,
	}

	mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
	pgxRows := pgxpoolmock.NewRows([]string{"token"}).ToPgxRows()
	mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(pgxRows, nil)

	repo := NewPostgresUserRepository(mockPool)
	_, err = repo.GetToken(token)

	// then
	assert.Error(t, err)
	//assert.NoError(t, validate.Struct(res))
	//assert.Equal(t, token, res)
}

//func Test(t *testing.T) {
//	tests := []struct {
//		name string
//		rows pgx.Rows
//	}{
//		{
//			name: "xxx",
//			rows:
//		},
//		// TODO: test cases
//	}
//	for _, test := range tests {
//		t.Run(test.name, func(t *testing.T) {
//			t.Parallel()
//			ctrl := gomock.NewController(t)
//			defer ctrl.Finish()
//
//			// given
//			jwtToken, err := jwt.GenerateToken("me")
//			require.NoError(t, err, "jwt token should be generated")
//
//			token := models.Token{
//				Token: jwtToken,
//			}
//
//			mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
//			expected := pgxpoolmock.NewRow(jwtToken)
//			mockPool.EXPECT().QueryRow(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(expected)
//
//			repo := NewPostgresUserRepository(mockPool)
//
//			res, err := repo.CreateToken(token)
//
//			// then
//			assert.NoError(t, err)
//			assert.NoError(t, validate.Struct(res))
//			assert.Equal(t, token, res)
//		})
//	}
//}
//
//func TestPostgresTokenRepository_GetToken1(t *testing.T) {
//	type fields struct {
//		client database.Client
//	}
//	type args struct {
//		token models.Token
//	}
//	tests := []struct {
//		name    string
//		fields  fields
//		args    args
//		want    models.Token
//		wantErr assert.ErrorAssertionFunc
//	}{
//		// TODO: Add test cases.
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			p := &PostgresTokenRepository{
//				client: tt.fields.client,
//			}
//			got, err := p.GetToken(tt.args.token)
//			if !tt.wantErr(t, err, fmt.Sprintf("GetToken(%v)", tt.args.token)) {
//				return
//			}
//			assert.Equalf(t, tt.want, got, "GetToken(%v)", tt.args.token)
//		})
//	}
//}

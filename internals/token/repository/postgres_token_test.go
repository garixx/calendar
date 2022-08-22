package repository

import (
	"calendar/internals/jwt"
	"calendar/internals/models"
	"calendar/internals/validate"
	"github.com/thanhpk/randstr"
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
	jwtToken, err := jwt.GenerateToken(randstr.String(5))
	require.NoError(t, err, "jwt token should be generated")

	type args struct {
		rows      pgx.Rows
		rowsError error
		token     models.Token
	}

	tests := []struct {
		name    string
		args    args
		want    models.Token
		wantErr bool
	}{
		{
			name: "positive case",
			args: args{
				rows:      pgxpoolmock.NewRows([]string{"token"}).AddRow(jwtToken).ToPgxRows(),
				rowsError: nil,
				token:     models.Token{Token: jwtToken},
			},
			want:    models.Token{Token: jwtToken},
			wantErr: false,
		},
		{
			name: "db request error",
			args: args{
				rows:      nil,
				rowsError: pgx.ErrNoRows,
				token:     models.Token{Token: jwtToken},
			},
			want:    models.Token{},
			wantErr: true,
		},
		{
			name: "not exist in DB",
			args: args{
				rows:      pgxpoolmock.NewRows([]string{"token"}).ToPgxRows(),
				rowsError: nil,
				token:     models.Token{Token: jwtToken},
			},
			want:    models.Token{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			mockPool := pgxpoolmock.NewMockPgxIface(ctrl)
			mockPool.EXPECT().Query(gomock.Any(), gomock.Any(), gomock.Eq(jwtToken)).Return(tt.args.rows, tt.args.rowsError)

			repo := NewPostgresUserRepository(mockPool)
			res, err := repo.GetToken(tt.args.token)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.NoError(t, validate.Struct(res))
				assert.Equal(t, tt.args.token, res)
			}
		})
	}
}

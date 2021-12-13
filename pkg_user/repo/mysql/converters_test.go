package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DarkSoul94/helpdesk2/models"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepo_toModelUser(t *testing.T) {
	r := NewRepo(&sql.DB{})
	tests := []struct {
		name  string
		input dbUser
		want  *models.User
	}{
		{
			name: "correct convert",
			input: dbUser{
				ID:      1,
				Name:    "test",
				Email:   "test@gmail.com",
				GroupID: 1,
				Department: sql.NullString{
					String: "test",
					Valid:  true,
				},
			},
			want: &models.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
				Group: &models.Group{
					ID: 1,
				},
				Department: "test",
			},
		},
	}
	for _, test := range tests {
		actual := r.toModelUser(test.input)
		assert.Equal(t, test.want, actual, test.name)
	}
}

func TestRepo_toDbUser(t *testing.T) {
	type fields struct {
		db *sqlx.DB
	}
	type args struct {
		user *models.User
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   dbUser
	}{
		{
			name:   "correct convert",
			fields: fields{db: &sqlx.DB{}},
			args: args{
				&models.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
					Group: &models.Group{
						ID: 1,
					},
					Department: "test",
				},
			},
			want: dbUser{
				ID:      1,
				Name:    "test",
				Email:   "test@gmail.com",
				GroupID: 1,
				Department: sql.NullString{
					String: "test",
					Valid:  true,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repo{
				db: tt.fields.db,
			}
			if got := r.toDbUser(tt.args.user); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repo.toDbUser() = %v, want %v", got, tt.want)
			}
		})
	}
}

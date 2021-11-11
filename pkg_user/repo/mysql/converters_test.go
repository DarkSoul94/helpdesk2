package mysql

import (
	"database/sql"
	"reflect"
	"testing"

	"github.com/DarkSoul94/helpdesk2/pkg_user"
	"github.com/DarkSoul94/helpdesk2/pkg_user/group_manager"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
)

func TestRepo_toModelUser(t *testing.T) {
	r := NewRepo(&sql.DB{})
	tests := []struct {
		name  string
		input dbUser
		want  *pkg_user.User
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
			want: &pkg_user.User{
				ID:    1,
				Name:  "test",
				Email: "test@gmail.com",
				Group: &group_manager.Group{
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
		user *pkg_user.User
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
				&pkg_user.User{
					ID:    1,
					Name:  "test",
					Email: "test@gmail.com",
					Group: &group_manager.Group{
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

package repository

import (
	"database/sql"
	"errors"
	"log"
	"testing"
	"tidy/pkg/model"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestUser_Create(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserStorage(db)

	type args struct {
		user model.User
	}
	type mockBehavior func(args args, id int64)

	testTable := []struct {
		name         string
		mockBehavior mockBehavior
		args         args
		id           int64
		wantErr      bool
	}{
		{
			name: "OK",
			args: args{

				user: model.User{
					Name:    "test",
					Surname: "test1",
				},
			},

			id: 22,
			mockBehavior: func(args args, id int64) {
				mock.ExpectExec("^INSERT").WithArgs(args.user.Name, args.user.Surname).WillReturnResult(sqlmock.NewResult(22, 1)).WillReturnError(nil)
			},
		},
		{
			name: "Empty",
			args: args{
				user: model.User{},
			},
			id: 0,
			mockBehavior: func(args args, id int64) {
				mock.ExpectExec("^INSERT").WithArgs(args.user.Name, args.user.Surname).WillReturnResult(sqlmock.NewResult(0, 0)).WillReturnError(errors.New("insert error"))
			},
			wantErr: true,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.mockBehavior(testCase.args, testCase.id)

			got, err := r.Create(&testCase.args.user)
			if testCase.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, testCase.id, got)
			}
		})
	}
}
func TestUser_Read(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserStorage(db)

	type args struct {
		userId int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		want    model.User
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "sername"}).AddRow(1, "title", "derbes")

				mock.ExpectQuery("^SELECT").
					WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				userId: 1,
			},
			want: model.User{1, "title", "derbes"},
		},
		{
			name: "Not Found",
			mock: func() {
				rows := sqlmock.NewRows([]string{"id", "name", "sername"})

				mock.ExpectQuery("^SELECT").
					WithArgs(1).WillReturnRows(rows)
			},
			input: args{
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()

			got, err := r.ReadOne(tt.input.userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUser_Delete(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserStorage(db)

	type args struct {
		userId int
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectExec("^DELETE").
					WithArgs(1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				userId: 1,
			},
		},
		{
			name: "Not Found",
			mock: func() {
				mock.ExpectExec("^DELETE").
					WithArgs(1).WillReturnError(sql.ErrNoRows)
			},
			input: args{
				userId: 1,
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := r.Delete(tt.input.userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}
}

func TestUser_Update(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	r := NewUserStorage(db)

	type args struct {
		userId int
		input  model.UpdateU
	}

	tests := []struct {
		name    string
		mock    func()
		input   args
		wantErr bool
	}{
		{
			name: "OK",
			mock: func() {
				mock.ExpectExec("^UPDATE").
					WithArgs("newname", "newsername", 1).WillReturnResult(sqlmock.NewResult(0, 1))
			},
			input: args{
				userId: 1,
				input: model.UpdateU{
					Name:    "newname",
					Surname: "newsername",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.mock()
			err := r.Update(&tt.input.input, tt.input.userId)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}
			assert.NoError(t, mock.ExpectationsWereMet())
		})
	}

}

package store

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"sync"

	"auth-app/model"

	"github.com/jinzhu/copier"
)

// ErrAlreadyExists is returned when a record with the same phone already exists in the store
var ErrAlreadyExists = errors.New("record already exists")

// UserStore is an interface to store user
type UserStore interface {
	// Save saves the user to the store
	Save(user *model.User) error
	// Find finds a user by phone
	Find(phone string) (*model.User, error)
}

// InMemoryUserStore stores user in memory
type InMemoryUserStore struct {
	mutex    sync.RWMutex
	data     map[string]*model.User
	dataPath string
}

// NewInMemoryUserStore returns a new InMemoryUserStore
func NewInMemoryUserStore(dataPath string) *InMemoryUserStore {
	return &InMemoryUserStore{
		data:     make(map[string]*model.User),
		dataPath: dataPath,
	}
}

func (store *InMemoryUserStore) PopulateDataFromFile() error {
	_, err := os.Stat(store.dataPath)
	if err == nil {
		users := []*model.User{}

		usersFromFile, err := store.readFromFile()
		if err != nil {
			return fmt.Errorf("failed to write user to file: %v", err.Error())
		}

		users = append(users, usersFromFile...)

		for _, user := range users {
			store.data[user.Phone] = user
		}
	}

	return nil
}

func (store *InMemoryUserStore) writeToFile(user *model.User) error {
	users := []*model.User{}

	_, err := os.Stat(store.dataPath)
	if err != nil {
		if os.IsNotExist(err) {
			users = append(users, user)
		} else {
			return fmt.Errorf("failed to write user to file: %v", err.Error())
		}
	} else {
		usersFromFile, err := store.readFromFile()
		if err != nil {
			return fmt.Errorf("failed to write user to file: %v", err.Error())
		}

		users = append(users, usersFromFile...)
		users = append(users, user)
	}

	data, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal user: %w", err)
	}

	err = ioutil.WriteFile(store.dataPath, data, 0644)
	if err != nil {
		return fmt.Errorf("failed to write data to file: %w", err)
	}

	return nil
}

func (store *InMemoryUserStore) readFromFile() ([]*model.User, error) {
	var users []*model.User

	data, err := ioutil.ReadFile(store.dataPath)
	if err != nil {
		return nil, fmt.Errorf("cannot read data from file: %w", err)
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil, fmt.Errorf("cannot unmarshal users: %w", err)
	}

	return users, nil
}

// Save saves the user to the store
func (store *InMemoryUserStore) Save(user *model.User) error {
	store.mutex.Lock()
	defer store.mutex.Unlock()

	if store.data[user.Phone] != nil {
		return ErrAlreadyExists
	}

	other, err := deepCopy(user)
	if err != nil {
		return err
	}

	store.data[other.Phone] = other

	// write to binary file
	err = store.writeToFile(user)
	if err != nil {
		return err
	}

	return nil
}

// Find finds a user by phone
func (store *InMemoryUserStore) Find(phone string) (*model.User, error) {
	store.mutex.RLock()
	defer store.mutex.RUnlock()

	user := store.data[phone]
	if user == nil {
		return nil, nil
	}

	return deepCopy(user)
}

func deepCopy(user *model.User) (*model.User, error) {
	other := &model.User{}

	err := copier.Copy(other, user)
	if err != nil {
		return nil, fmt.Errorf("cannot copy user data: %w", err)
	}

	return other, nil
}

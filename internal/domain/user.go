package domain

import (
	"fmt"
	"sort"
)

type User struct {
	id      int    `json:"id"`
	name    string `json:"name"`
	age     uint8  `json:"age"`
	friends []int  `json:"friends,omitempty"`
}

func NewUser(id int, name string, age uint8, friends []int) User {
	return User{id: id, name: name, age: age, friends: friends}
}

func (user *User) SetId(id int) {
	user.id = id
}

func (user *User) SetName(name string) {
	user.name = name
}

func (user *User) SetAge(age uint8) {
	user.age = age
}

func (user *User) ID() int {
	return user.id
}

func (user *User) Name() string {
	return user.name
}

func (user *User) Age() uint8 {
	return user.age
}

func (user *User) Friends() []int {
	return user.friends
}

func (user *User) SetFriend(userId int) error {
	bFriend, index := user.isFriend(userId)
	if bFriend {
		return fmt.Errorf("domain (id %d) in friends", userId)
	}
	user.friends = insertAt(user.friends, index, userId)
	return nil
}

func (user *User) SetFriends(friendsId []int) error {
	for _, userId := range friendsId {
		err := user.SetFriend(userId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (user *User) isFriend(userId int) (bool, int) {
	index := sort.SearchInts(user.friends, userId)
	if index == len(user.friends) {
		return false, index
	}
	return true, index
}

func (user *User) DeleteFriend(userId int) (int, error) {
	bFriend, index := user.isFriend(userId)
	if !bFriend {
		return userId, fmt.Errorf("domain (id %d) in not friend for domain (id %d)", userId, user.id)
	}
	copy(user.friends[index:], user.friends[index+1:])
	user.friends[len(user.friends)-1] = -1
	user.friends = user.friends[:len(user.friends)-1]
	return userId, nil
}

func insertAt(data []int, i int, v int) []int {
	if i == len(data) {
		return append(data, v)
	}
	data = append(data[:i+1], data[i:]...)
	data[i] = v
	return data
}

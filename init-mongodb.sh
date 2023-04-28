#!/bin/bash

mongosh <<EOF
use amaranth

db.createCollection("users")
db.users.insertOne({
  "_id": ObjectId("61423c79c7f5467890123456"),
  "name": "John",
  "lastname": "Doe",
  "email": "johndoe@example.com",
  "password": "<hashed-password>",
  "creationDate": "2022-09-15T10:30:00Z",
  "status": "active",
  "messages": [
    {
      "_id": ObjectId("61423c79c7f5467890abcdef"),
      "sender": "Alice",
      "recipient": "John",
      "content": "Hello John, how are you?",
      "timestamp": "2022-09-15T11:00:00Z"
    },
    {
      "_id": ObjectId("61423c79c7f5467890fedcba"),
      "sender": "John",
      "recipient": "Alice",
      "content": "Hi Alice, I'm doing well, thank you!",
      "timestamp": "2022-09-15T11:05:00Z"
    }
  ],
  "channels": [
    {
      "_id": ObjectId("61423c79c7f5467890c1c1c1"),
      "name": "general",
      "description": "General channel for discussions",
      "createdBy": "John",
      "creationDate": "2022-09-15T12:00:00Z",
      "messages": [
        {
          "_id": ObjectId("61423c79c7f5467890d3d3d3"),
          "sender": "Bob",
          "recipient": "John",
          "content": "Hey John, have you seen the latest update?",
          "timestamp": "2022-09-15T12:30:00Z"
        }
      ]
    },
    {
      "_id": ObjectId("61423c79c7f5467890d2d2d2"),
      "name": "team-1",
      "description": "Channel for Team 1 members",
      "createdBy": "Alice",
      "creationDate": "2022-09-15T12:30:00Z",
      "messages": [
        {
          "_id": ObjectId("61423c79c7f5467890e4e4e4"),
          "sender": "Alice",
          "recipient": "John",
          "content": "John, we have a meeting tomorrow at 2 PM.",
          "timestamp": "2022-09-15T13:00:00Z"
        },
        {
          "_id": ObjectId("61423c79c7f5467890f5f5f5"),
          "sender": "Dave",
          "recipient": "John",
          "content": "Hey John, can you review the latest pull request?",
          "timestamp": "2022-09-15T14:00:00Z"
        }
      ]
    }
  ]
})

if [ -n "$MONGODB_USERNAME" ] && [ -n "$MONGODB_PASSWORD" ]; then
  db.createUser({
    user: "$MONGODB_USERNAME",
    pwd: "$MONGODB_PASSWORD",
    roles: ["root"]
  })
fi
EOF

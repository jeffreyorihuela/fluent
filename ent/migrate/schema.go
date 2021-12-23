// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// ChatsColumns holds the columns for the "chats" table.
	ChatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "type", Type: field.TypeEnum, Enums: []string{"private", "public"}},
		{Name: "deleted", Type: field.TypeBool, Default: false},
	}
	// ChatsTable holds the schema information for the "chats" table.
	ChatsTable = &schema.Table{
		Name:       "chats",
		Columns:    ChatsColumns,
		PrimaryKey: []*schema.Column{ChatsColumns[0]},
	}
	// MessagesColumns holds the columns for the "messages" table.
	MessagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "body", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "user_messages", Type: field.TypeInt, Nullable: true},
	}
	// MessagesTable holds the schema information for the "messages" table.
	MessagesTable = &schema.Table{
		Name:       "messages",
		Columns:    MessagesColumns,
		PrimaryKey: []*schema.Column{MessagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "messages_users_messages",
				Columns:    []*schema.Column{MessagesColumns[3]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "username", Type: field.TypeString, Default: "unknown"},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "password", Type: field.TypeString, Default: "isnotasecret"},
		{Name: "private_key", Type: field.TypeString, Default: "unknown"},
		{Name: "public_key", Type: field.TypeBytes},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// ChatMembersColumns holds the columns for the "chat_members" table.
	ChatMembersColumns = []*schema.Column{
		{Name: "chat_id", Type: field.TypeInt},
		{Name: "user_id", Type: field.TypeInt},
	}
	// ChatMembersTable holds the schema information for the "chat_members" table.
	ChatMembersTable = &schema.Table{
		Name:       "chat_members",
		Columns:    ChatMembersColumns,
		PrimaryKey: []*schema.Column{ChatMembersColumns[0], ChatMembersColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chat_members_chat_id",
				Columns:    []*schema.Column{ChatMembersColumns[0]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "chat_members_user_id",
				Columns:    []*schema.Column{ChatMembersColumns[1]},
				RefColumns: []*schema.Column{UsersColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ChatMessagesColumns holds the columns for the "chat_messages" table.
	ChatMessagesColumns = []*schema.Column{
		{Name: "chat_id", Type: field.TypeInt},
		{Name: "message_id", Type: field.TypeInt},
	}
	// ChatMessagesTable holds the schema information for the "chat_messages" table.
	ChatMessagesTable = &schema.Table{
		Name:       "chat_messages",
		Columns:    ChatMessagesColumns,
		PrimaryKey: []*schema.Column{ChatMessagesColumns[0], ChatMessagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "chat_messages_chat_id",
				Columns:    []*schema.Column{ChatMessagesColumns[0]},
				RefColumns: []*schema.Column{ChatsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "chat_messages_message_id",
				Columns:    []*schema.Column{ChatMessagesColumns[1]},
				RefColumns: []*schema.Column{MessagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		ChatsTable,
		MessagesTable,
		UsersTable,
		ChatMembersTable,
		ChatMessagesTable,
	}
)

func init() {
	MessagesTable.ForeignKeys[0].RefTable = UsersTable
	ChatMembersTable.ForeignKeys[0].RefTable = ChatsTable
	ChatMembersTable.ForeignKeys[1].RefTable = UsersTable
	ChatMessagesTable.ForeignKeys[0].RefTable = ChatsTable
	ChatMessagesTable.ForeignKeys[1].RefTable = MessagesTable
}

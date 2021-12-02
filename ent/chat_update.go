// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/amaru0601/fluent/ent/chat"
	"github.com/amaru0601/fluent/ent/message"
	"github.com/amaru0601/fluent/ent/predicate"
	"github.com/amaru0601/fluent/ent/user"
)

// ChatUpdate is the builder for updating Chat entities.
type ChatUpdate struct {
	config
	hooks    []Hook
	mutation *ChatMutation
}

// Where appends a list predicates to the ChatUpdate builder.
func (cu *ChatUpdate) Where(ps ...predicate.Chat) *ChatUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetName sets the "name" field.
func (cu *ChatUpdate) SetName(s string) *ChatUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetType sets the "type" field.
func (cu *ChatUpdate) SetType(c chat.Type) *ChatUpdate {
	cu.mutation.SetType(c)
	return cu
}

// SetDeleted sets the "deleted" field.
func (cu *ChatUpdate) SetDeleted(b bool) *ChatUpdate {
	cu.mutation.SetDeleted(b)
	return cu
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (cu *ChatUpdate) SetNillableDeleted(b *bool) *ChatUpdate {
	if b != nil {
		cu.SetDeleted(*b)
	}
	return cu
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cu *ChatUpdate) AddMemberIDs(ids ...int) *ChatUpdate {
	cu.mutation.AddMemberIDs(ids...)
	return cu
}

// AddMembers adds the "members" edges to the User entity.
func (cu *ChatUpdate) AddMembers(u ...*User) *ChatUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.AddMemberIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cu *ChatUpdate) AddMessageIDs(ids ...int) *ChatUpdate {
	cu.mutation.AddMessageIDs(ids...)
	return cu
}

// AddMessages adds the "messages" edges to the Message entity.
func (cu *ChatUpdate) AddMessages(m ...*Message) *ChatUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.AddMessageIDs(ids...)
}

// Mutation returns the ChatMutation object of the builder.
func (cu *ChatUpdate) Mutation() *ChatMutation {
	return cu.mutation
}

// ClearMembers clears all "members" edges to the User entity.
func (cu *ChatUpdate) ClearMembers() *ChatUpdate {
	cu.mutation.ClearMembers()
	return cu
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cu *ChatUpdate) RemoveMemberIDs(ids ...int) *ChatUpdate {
	cu.mutation.RemoveMemberIDs(ids...)
	return cu
}

// RemoveMembers removes "members" edges to User entities.
func (cu *ChatUpdate) RemoveMembers(u ...*User) *ChatUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cu.RemoveMemberIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cu *ChatUpdate) ClearMessages() *ChatUpdate {
	cu.mutation.ClearMessages()
	return cu
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cu *ChatUpdate) RemoveMessageIDs(ids ...int) *ChatUpdate {
	cu.mutation.RemoveMessageIDs(ids...)
	return cu
}

// RemoveMessages removes "messages" edges to Message entities.
func (cu *ChatUpdate) RemoveMessages(m ...*Message) *ChatUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cu.RemoveMessageIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(cu.hooks) == 0 {
		if err = cu.check(); err != nil {
			return 0, err
		}
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cu.check(); err != nil {
				return 0, err
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChatUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChatUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChatUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ChatUpdate) check() error {
	if v, ok := cu.mutation.GetType(); ok {
		if err := chat.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (cu *ChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chat.Table,
			Columns: chat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chat.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chat.FieldName,
		})
	}
	if value, ok := cu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chat.FieldType,
		})
	}
	if value, ok := cu.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chat.FieldDeleted,
		})
	}
	if cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cu.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cu.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChatUpdateOne is the builder for updating a single Chat entity.
type ChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChatMutation
}

// SetName sets the "name" field.
func (cuo *ChatUpdateOne) SetName(s string) *ChatUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetType sets the "type" field.
func (cuo *ChatUpdateOne) SetType(c chat.Type) *ChatUpdateOne {
	cuo.mutation.SetType(c)
	return cuo
}

// SetDeleted sets the "deleted" field.
func (cuo *ChatUpdateOne) SetDeleted(b bool) *ChatUpdateOne {
	cuo.mutation.SetDeleted(b)
	return cuo
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (cuo *ChatUpdateOne) SetNillableDeleted(b *bool) *ChatUpdateOne {
	if b != nil {
		cuo.SetDeleted(*b)
	}
	return cuo
}

// AddMemberIDs adds the "members" edge to the User entity by IDs.
func (cuo *ChatUpdateOne) AddMemberIDs(ids ...int) *ChatUpdateOne {
	cuo.mutation.AddMemberIDs(ids...)
	return cuo
}

// AddMembers adds the "members" edges to the User entity.
func (cuo *ChatUpdateOne) AddMembers(u ...*User) *ChatUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.AddMemberIDs(ids...)
}

// AddMessageIDs adds the "messages" edge to the Message entity by IDs.
func (cuo *ChatUpdateOne) AddMessageIDs(ids ...int) *ChatUpdateOne {
	cuo.mutation.AddMessageIDs(ids...)
	return cuo
}

// AddMessages adds the "messages" edges to the Message entity.
func (cuo *ChatUpdateOne) AddMessages(m ...*Message) *ChatUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.AddMessageIDs(ids...)
}

// Mutation returns the ChatMutation object of the builder.
func (cuo *ChatUpdateOne) Mutation() *ChatMutation {
	return cuo.mutation
}

// ClearMembers clears all "members" edges to the User entity.
func (cuo *ChatUpdateOne) ClearMembers() *ChatUpdateOne {
	cuo.mutation.ClearMembers()
	return cuo
}

// RemoveMemberIDs removes the "members" edge to User entities by IDs.
func (cuo *ChatUpdateOne) RemoveMemberIDs(ids ...int) *ChatUpdateOne {
	cuo.mutation.RemoveMemberIDs(ids...)
	return cuo
}

// RemoveMembers removes "members" edges to User entities.
func (cuo *ChatUpdateOne) RemoveMembers(u ...*User) *ChatUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return cuo.RemoveMemberIDs(ids...)
}

// ClearMessages clears all "messages" edges to the Message entity.
func (cuo *ChatUpdateOne) ClearMessages() *ChatUpdateOne {
	cuo.mutation.ClearMessages()
	return cuo
}

// RemoveMessageIDs removes the "messages" edge to Message entities by IDs.
func (cuo *ChatUpdateOne) RemoveMessageIDs(ids ...int) *ChatUpdateOne {
	cuo.mutation.RemoveMessageIDs(ids...)
	return cuo
}

// RemoveMessages removes "messages" edges to Message entities.
func (cuo *ChatUpdateOne) RemoveMessages(m ...*Message) *ChatUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return cuo.RemoveMessageIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChatUpdateOne) Select(field string, fields ...string) *ChatUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Chat entity.
func (cuo *ChatUpdateOne) Save(ctx context.Context) (*Chat, error) {
	var (
		err  error
		node *Chat
	)
	if len(cuo.hooks) == 0 {
		if err = cuo.check(); err != nil {
			return nil, err
		}
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = cuo.check(); err != nil {
				return nil, err
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChatUpdateOne) SaveX(ctx context.Context) *Chat {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChatUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChatUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ChatUpdateOne) check() error {
	if v, ok := cuo.mutation.GetType(); ok {
		if err := chat.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf("ent: validator failed for field \"type\": %w", err)}
		}
	}
	return nil
}

func (cuo *ChatUpdateOne) sqlSave(ctx context.Context) (_node *Chat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   chat.Table,
			Columns: chat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: chat.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Chat.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, chat.FieldID)
		for _, f := range fields {
			if !chat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != chat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: chat.FieldName,
		})
	}
	if value, ok := cuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeEnum,
			Value:  value,
			Column: chat.FieldType,
		})
	}
	if value, ok := cuo.mutation.Deleted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: chat.FieldDeleted,
		})
	}
	if cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMembersIDs(); len(nodes) > 0 && !cuo.mutation.MembersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MembersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MembersTable,
			Columns: chat.MembersPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedMessagesIDs(); len(nodes) > 0 && !cuo.mutation.MessagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.MessagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   chat.MessagesTable,
			Columns: chat.MessagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Chat{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{chat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}

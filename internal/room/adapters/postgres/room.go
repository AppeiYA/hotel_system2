package room_postgres

import (
	"context"
	"database/sql"
	"errors"
	"hotel_system2/internal/room/domain"
	"time"
)

func (r *Repository) Create(ctx context.Context, room *domain.Room) error {
	row := roomRowFromDomain(room)
	if err := r.executor(ctx).GetContext(ctx, &row, CREATE_ROOM,
		row.RoomNumber, row.Type, row.Rate, row.Status); err != nil {
		return err
	}
	*room = *row.toDomain()
	return nil
}

func (r *Repository) FindByID(ctx context.Context, id string) (*domain.Room, error) {
	var row roomRow
	err := r.executor(ctx).GetContext(ctx, &row, GET_ROOM_BY_ID, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrRoomNotFound
	}
	if err != nil {
		return nil, err
	}
	room := row.toDomain()
	return room, nil
}

func (r *Repository) FindByNumber(ctx context.Context, roomNumber string) (*domain.Room, error) {
	var row roomRow
	err := r.executor(ctx).GetContext(ctx, &row, GET_ROOM_BY_NUMBER, roomNumber)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	room := row.toDomain()
	return room, nil
}

func (r *Repository) List(ctx context.Context) ([]*domain.Room, error) {
	var rows []roomRow
	if err := r.executor(ctx).SelectContext(ctx, &rows, LIST_ROOMS); err != nil {
		return nil, err
	}
	rooms := make([]*domain.Room, len(rows))
	for i, row := range rows {
		rooms[i] = row.toDomain()
	}
	return rooms, nil
}

func (r *Repository) Update(ctx context.Context, room *domain.Room) error {
	row := roomRowFromDomain(room)
	if err := r.executor(ctx).GetContext(ctx, &row, UPDATE_ROOM,
		row.RoomNumber, row.Type, row.Rate, row.Status, row.ID); err != nil {
		return err
	}
	*room = *row.toDomain()
	return nil
}

func (r *Repository) Delete(ctx context.Context, id string) error {
	_, err := r.executor(ctx).ExecContext(ctx, DELETE_ROOM, id)
	return err
}

func (r *Repository) UpdateStatus(ctx context.Context, id string, status domain.RoomStatus) error {
	exec := r.executor(ctx)

	var row roomRow
	if err := exec.GetContext(ctx, &row, GET_ROOM_BY_ID_FOR_UPDATE, id); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return errors.New("room not found")
		}
		return err
	}

	_, err := exec.ExecContext(ctx, UPDATE_ROOM_STATUS, status, id)
	return err
}

func (r *Repository) FindAvailable(
	ctx context.Context,
	roomType domain.RoomType,
	checkIn time.Time,
	checkOut time.Time,
) (*domain.Room, error) {
	var row roomRow
	err := r.executor(ctx).GetContext(ctx, &row, FIND_AVAILABLE_ROOM, roomType, checkIn, checkOut)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	room := row.toDomain()
	return room, nil
}

func (r *Repository) FindByIDForUpdate(ctx context.Context, id string) (*domain.Room, error) {
	var row roomRow
	err := r.executor(ctx).GetContext(ctx, &row, GET_ROOM_BY_ID_FOR_UPDATE, id)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, domain.ErrRoomNotFound
	}
	if err != nil {
		return nil, err
	}
	room := row.toDomain()
	return room, nil
}

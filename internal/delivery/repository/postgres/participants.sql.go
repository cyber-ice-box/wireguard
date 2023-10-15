// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0
// source: participants.sql

package postgres

import (
	"context"

	"github.com/google/uuid"
	"github.com/sqlc-dev/pqtype"
)

const getParticipantsInRunningEvents = `-- name: GetParticipantsInRunningEvents :many
select participants.id,
       participants.vpn_ip_address,
       participants.vpn_public_key,
       teams.lab_cidr,
       COALESCE(teams.lab_permitted, false)
from participants
         left join teams on teams.id = participants.team_id
         inner join events on events.id = participants.event_id
where events.withdraw_time > now()
`

type GetParticipantsInRunningEventsRow struct {
	ID           uuid.UUID   `json:"id"`
	VpnIpAddress pqtype.Inet `json:"vpn_ip_address"`
	VpnPublicKey string      `json:"vpn_public_key"`
	LabCidr      pqtype.Inet `json:"lab_cidr"`
	LabPermitted bool        `json:"lab_permitted"`
}

func (q *Queries) GetParticipantsInRunningEvents(ctx context.Context) ([]GetParticipantsInRunningEventsRow, error) {
	rows, err := q.query(ctx, q.getParticipantsInRunningEventsStmt, getParticipantsInRunningEvents)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []GetParticipantsInRunningEventsRow{}
	for rows.Next() {
		var i GetParticipantsInRunningEventsRow
		if err := rows.Scan(
			&i.ID,
			&i.VpnIpAddress,
			&i.VpnPublicKey,
			&i.LabCidr,
			&i.LabPermitted,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
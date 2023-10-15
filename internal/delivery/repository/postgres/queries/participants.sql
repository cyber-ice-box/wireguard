-- name: GetParticipantsInRunningEvents :many
select participants.id,
       participants.vpn_ip_address,
       participants.vpn_public_key,
       teams.lab_cidr,
       COALESCE(teams.lab_permitted, false)
from participants
         left join teams on teams.id = participants.team_id
         inner join events on events.id = participants.event_id
where events.withdraw_time > now();
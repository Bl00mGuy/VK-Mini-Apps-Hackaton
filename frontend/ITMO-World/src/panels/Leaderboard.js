import React, {useEffect, useState} from 'react';
import {Avatar, Cell, Group, Panel, PanelHeader} from '@vkontakte/vkui';
import PropTypes from 'prop-types';
import {fetchLeaderboard} from '../api';

const LeaderboardPanel = ({id}) => {
    const [leaderboard, setLeaderboard] = useState([]);

    useEffect(() => {
        const getLeaderboard = async () => {
            const data = await fetchLeaderboard();
            setLeaderboard(data);
        };
        getLeaderboard();
    }, []);

    return (
        <Panel id={id}>
            <PanelHeader>Рейтинг</PanelHeader>
            <Group>
                {leaderboard.map((user, index) => (
                    <Cell
                        key={user.id}
                        before={<Avatar src={user.avatar}/>}
                        description={`Очки: ${user.points}`}
                    >
                        {index + 1}. {user.name}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

LeaderboardPanel.propTypes = {
    id: PropTypes.string.isRequired,
};

export {LeaderboardPanel};

import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, Cell, ScreenSpinner } from '@vkontakte/vkui';
import { fetchLeaderboard } from '../api';
import PropTypes from 'prop-types';

const LeaderboardPanel = ({ id, goTo }) => {
    const [leaderboard, setLeaderboard] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getLeaderboard = async () => {
            try {
                const fetchedLeaderboard = await fetchLeaderboard();
                setLeaderboard(fetchedLeaderboard);
            } catch (err) {
                setError('Ошибка при загрузке лидерборда');
            } finally {
                setLoading(false);
            }
        };
        getLeaderboard();
    }, []);

    if (loading) return <ScreenSpinner size="large" />;
    if (error) return <div>{error}</div>;

    return (
        <Panel id={id}>
            <PanelHeader>Лидерборд</PanelHeader>
            <Group>
                {leaderboard.map((player) => (
                    <Cell key={player.id} onClick={() => goTo('home')}>
                        {player.name} - {player.score}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

LeaderboardPanel.propTypes = {
    id: PropTypes.string.isRequired,
    goTo: PropTypes.func.isRequired,
};

export { LeaderboardPanel };

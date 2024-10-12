import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, Cell, ScreenSpinner } from '@vkontakte/vkui';
import { fetchAchievements } from '../api';
import PropTypes from 'prop-types';

const AchievementsPanel = ({ id, goTo }) => {
    const [achievements, setAchievements] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getAchievements = async () => {
            try {
                const fetchedAchievements = await fetchAchievements();
                setAchievements(fetchedAchievements);
            } catch (err) {
                setError('Ошибка при загрузке достижений');
            } finally {
                setLoading(false);
            }
        };
        getAchievements();
    }, []);

    if (loading) return <ScreenSpinner size="large" />;
    if (error) return <div>{error}</div>;

    return (
        <Panel id={id}>
            <PanelHeader>Достижения</PanelHeader>
            <Group>
                {achievements.map((achievement) => (
                    <Cell key={achievement.id} onClick={() => goTo('home')}>
                        {achievement.title}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

AchievementsPanel.propTypes = {
    id: PropTypes.string.isRequired,
    goTo: PropTypes.func.isRequired,
};

export { AchievementsPanel };

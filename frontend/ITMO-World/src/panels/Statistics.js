import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, Cell, ScreenSpinner } from '@vkontakte/vkui';
import { fetchStatistics } from '../api';
import PropTypes from 'prop-types';

const StatisticsPanel = ({ id, goTo }) => {
    const [statistics, setStatistics] = useState([]);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);

    useEffect(() => {
        const getStatistics = async () => {
            try {
                const fetchedStatistics = await fetchStatistics();
                setStatistics(fetchedStatistics);
            } catch (err) {
                setError('Ошибка при загрузке статистики');
            } finally {
                setLoading(false);
            }
        };
        getStatistics();
    }, []);

    if (loading) return <ScreenSpinner size="large" />;
    if (error) return <div>{error}</div>;

    return (
        <Panel id={id}>
            <PanelHeader>Статистика</PanelHeader>
            <Group>
                {statistics.map((stat) => (
                    <Cell key={stat.id} onClick={() => goTo('home')}>
                        {stat.title}: {stat.value}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

StatisticsPanel.propTypes = {
    id: PropTypes.string.isRequired,
    goTo: PropTypes.func.isRequired,
};

export { StatisticsPanel };

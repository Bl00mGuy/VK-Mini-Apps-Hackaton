import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, SimpleCell } from '@vkontakte/vkui';
import PropTypes from 'prop-types';
import { fetchStatistics } from '../api';

const StatisticsPanel = ({ id }) => {
    const [statistics, setStatistics] = useState(null);

    useEffect(() => {
        const getStatistics = async () => {
            const data = await fetchStatistics();
            setStatistics(data);
        };
        getStatistics();
    }, []);

    return (
        <Panel id={id}>
            <PanelHeader>Статистика</PanelHeader>
            {statistics ? (
                <Group>
                    <SimpleCell>Завершено заданий: {statistics.completedTasks}</SimpleCell>
                    <SimpleCell>Получено достижений: {statistics.achievements}</SimpleCell>
                    <SimpleCell>Общий рейтинг: {statistics.rank}</SimpleCell>
                </Group>
            ) : (
                <Group>Загрузка данных...</Group>
            )}
        </Panel>
    );
};

StatisticsPanel.propTypes = {
    id: PropTypes.string.isRequired,
};

export {StatisticsPanel};

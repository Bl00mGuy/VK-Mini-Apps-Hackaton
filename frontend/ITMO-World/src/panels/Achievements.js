import React, {useEffect, useState} from 'react';
import {Cell, Group, Panel, PanelHeader} from '@vkontakte/vkui';
import PropTypes from 'prop-types';
import {fetchAchievements} from '../api';

const AchievementsPanel = ({id}) => {
    const [achievements, setAchievements] = useState([]);

    useEffect(() => {
        const getAchievements = async () => {
            const fetchedAchievements = await fetchAchievements();
            setAchievements(fetchedAchievements);
        };
        getAchievements();
    }, []);

    return (
        <Panel id={id}>
            <PanelHeader>Достижения</PanelHeader>
            <Group>
                {achievements.map((achievement) => (
                    <Cell key={achievement.id} description={achievement.description}>
                        {achievement.title}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

AchievementsPanel.propTypes = {
    id: PropTypes.string.isRequired,
};

export {AchievementsPanel};

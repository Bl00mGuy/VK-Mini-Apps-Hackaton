import React from 'react';
import { Panel, PanelHeader, Group, Button } from '@vkontakte/vkui';
import PropTypes from 'prop-types';

const OnboardingPanel = ({ id, goTo }) => {
    return (
        <Panel id={id}>
            <PanelHeader>Добро пожаловать</PanelHeader>
            <Group>
                <p>Познакомьтесь с возможностями нашего приложения!</p>
                <Button size="l" stretched onClick={() => goTo('home')}>
                    Начать
                </Button>
            </Group>
        </Panel>
    );
};

OnboardingPanel.propTypes = {
    id: PropTypes.string.isRequired,
    goTo: PropTypes.func.isRequired,
};

export default OnboardingPanel;

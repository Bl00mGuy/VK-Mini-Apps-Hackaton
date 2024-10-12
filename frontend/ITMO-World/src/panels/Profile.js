import React from 'react';
import { Panel, PanelHeader, Placeholder } from '@vkontakte/vkui';
import PropTypes from 'prop-types';

const ProfilePanel = ({ id }) => {
    return (
        <Panel id={id}>
            <PanelHeader>Профиль</PanelHeader>
            <Placeholder>Информация о пользователе</Placeholder>
        </Panel>
    );
};

ProfilePanel.propTypes = {
    id: PropTypes.string.isRequired,
};

export default ProfilePanel;

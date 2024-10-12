import React from 'react';
import {Avatar, Button, Cell, Group, Panel, PanelHeader} from '@vkontakte/vkui';
import PropTypes from 'prop-types';

const HomePanel = ({id, fetchedUser, goTo}) => {
    return (
        <Panel id={id}>
            <PanelHeader>Главная</PanelHeader>
            {fetchedUser && (
                <Group>
                    <Cell
                        before={fetchedUser.photo_200 && <Avatar src={fetchedUser.photo_200}/>}
                        description={fetchedUser.city?.title}
                    >
                        {`${fetchedUser.first_name} ${fetchedUser.last_name}`}
                    </Cell>
                </Group>
            )}
            <Group>
                <Button size="l" stretched onClick={() => goTo('profile')}>
                    Профиль
                </Button>
                <Button size="l" stretched onClick={() => goTo('tasks')}>
                    Задания
                </Button>
                <Button size="l" stretched onClick={() => goTo('achievements')}>
                    Достижения
                </Button>
                <Button size="l" stretched onClick={() => goTo('leaderboard')}>
                    Лидерборд
                </Button>
            </Group>
        </Panel>
    );
};

HomePanel.propTypes = {
    id: PropTypes.string.isRequired,
    fetchedUser: PropTypes.object,
    goTo: PropTypes.func.isRequired,
};

export {HomePanel};

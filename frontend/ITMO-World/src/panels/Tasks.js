import React, { useState, useEffect } from 'react';
import { Panel, PanelHeader, Group, Cell, Button } from '@vkontakte/vkui';
import PropTypes from 'prop-types';
import { fetchTasks, completeTask } from '../api';

const TasksPanel = ({ id }) => {
    const [tasks, setTasks] = useState([]);

    useEffect(() => {
        const getTasks = async () => {
            const fetchedTasks = await fetchTasks();
            setTasks(fetchedTasks);
        };
        getTasks();
    }, []);

    const handleCompleteTask = async (taskId) => {
        await completeTask(taskId);
        setTasks(tasks.map(task =>
            task.id === taskId ? { ...task, completed: true } : task
        ));
    };

    return (
        <Panel id={id}>
            <PanelHeader>Задания</PanelHeader>
            <Group>
                {tasks.map((task) => (
                    <Cell key={task.id} description={task.description} asideContent={
                        <Button size="m" onClick={() => handleCompleteTask(task.id)} disabled={task.completed}>
                            {task.completed ? 'Выполнено' : 'Выполнить'}
                        </Button>
                    }>
                        {task.title}
                    </Cell>
                ))}
            </Group>
        </Panel>
    );
};

TasksPanel.propTypes = {
    id: PropTypes.string.isRequired,
};

export { TasksPanel };
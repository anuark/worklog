import React from 'react';
import { CardActions } from 'material-ui/Card';
import { CreateButton, RefreshButton } from 'admin-on-rest';
import GenerateButton from './GenerateButton';

const cardActionStyle = {
    zIndex: 2,
    display: 'inline-block',
    float: 'right',
};

const TaskListActions = ({ basePath, data }) => (
    <CardActions style={cardActionStyle}>
        <GenerateButton record={data} />
        <CreateButton basePath={basePath} />
        <RefreshButton />
    </CardActions>
);

export default TaskListActions;

import React from 'react';
import { List, Datagrid, Edit, Create, SimpleForm, DateField, TextField, EditButton, DisabledInput, TextInput, /*LongTextInput, DateInput*/ } from 'admin-on-rest';
import BookIcon from 'material-ui/svg-icons/action/book';
export const TaskIcon = BookIcon;

export const TaskList = (props) => (
    <List {...props} sort={{field: 'created', order: 'DESC'}}>
        <Datagrid>
            <TextField source="description" />
            <DateField source="created" />
            <EditButton basePath="/tasks" />
        </Datagrid>
    </List>
);

const TaskTitle = ({ record }) => {
    return <span>Task {record ? `"${record.title}"` : ''}</span>;
};

export const TaskEdit = (props) => (
    <Edit title={<TaskTitle />} {...props}>
        <SimpleForm>
            <DisabledInput source="id" />
            <TextInput source="description" />
            <DisabledInput source="created" />
        </SimpleForm>
    </Edit>
);

export const TaskCreate = (props) => (
    <Create title="Create a Task" {...props}>
        <SimpleForm>
            <TextInput source="description" />
            <DisabledInput source="created" />
        </SimpleForm>
    </Create>
);
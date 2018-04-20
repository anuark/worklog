import React from 'react';
import BookIcon from 'material-ui/svg-icons/action/book';
import {List, Create, DataGrid, TextField, DateField, TextInput, DateInput, SimpleForm} from 'admin-on-rest';

export const InvoiceIcon = BookIcon;

export const InvoiceList = (props) => (
    <List {...props} sort={{field: 'created', order: 'DESC'}} title="Invoice list">
        <DataGrid>
            <TextField source="name" />
            <DateField source="startDate" />
            <DateField source="endDate" />
        </DataGrid>
    </List>
);

export const InvoiceCreate = (props) => (
    <Create title="Create invoice" {...props}>
        <SimpleForm>
            <TextInput source="name"/>
            <DateInput source="startDate"/>
            <DateInput source="endDate"/>
        </SimpleForm>
    </Create>
);

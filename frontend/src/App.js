// in app.js
import React from 'react';
import { jsonServerRestClient, Admin, Resource, fetchUtils } from 'admin-on-rest';
import { TaskList, TaskEdit, TaskCreate, TaskIcon } from './tasks';
import { InvoiceList, InvoiceCreate, InvoiceIcon } from './invoices';
import authClient from './AuthClient';
// import MyLogin from './Login';
import MyLogin from './Login';
// import MyLogoutButton from './MyLogoutButton'

const httpClient = (url, options = {}) => {
    if (!options.headers) {
        options.headers = new Headers({ Accept: 'application/json' })
    }

    options.headers.set('Authorization', 'Bearer ' + localStorage.getItem('token'));
    options.noCors = true;
    return fetchUtils.fetchJson(url, options);
}

const restClient = jsonServerRestClient('http://localhost:8000', httpClient);

const App = () => (
    <Admin loginPage={MyLogin} /*logoutButton={MyLogoutButton}*/ authClient={authClient} restClient={restClient} title="Worklog Dashboard">
        <Resource name="tasks" list={TaskList} edit={TaskEdit} create={TaskCreate} icon={TaskIcon}/>
        <Resource name="invoices" list={InvoiceList} create={InvoiceCreate} icon={InvoiceIcon}/>
    </Admin>
);

export default App;
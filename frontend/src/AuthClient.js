import { AUTH_LOGIN, AUTH_LOGOUT, AUTH_ERROR, AUTH_CHECK } from 'admin-on-rest'

export default (type, params) => {
    // called on login submit
    if (type === AUTH_LOGIN) {
        const { email, password } = params;
        
        const request = new Request("http://localhost:8000/auth", {
            method: 'POST',
            body: JSON.stringify({ email, password }),
            // headers: new Headers({ 'Content-Type': 'application/x-www-form-urlencoded' })
            headers: new Headers({ 'Content-Type': 'application/json' })
        });

        return fetch(request)
            .then(response => {
                if (response.status === 400) {
                    throw new Error("Username or password is invalid.");
                } else if (response.status < 200 || response.status >= 300) {
                    throw new Error(response.statusText);
                }
                return response.json();
            })
            .then(({ token }) => {
                localStorage.setItem('token', token)
            });
    }
    // called when the user clicks on the logout button
    if (type === AUTH_LOGOUT) {
        localStorage.removeItem('token');
        return Promise.resolve();
    }

    if (type === AUTH_ERROR) {
        const { status } = params;
        if (status === 401 || status === 403) {
            localStorage.removeItem('token');
            return Promise.reject();
        }
        Promise.resolve();
    }

    // called when the user navigates to a new location
    if (type === AUTH_CHECK) {
        return localStorage.getItem('token') ? Promise.resolve() : Promise.reject();
    }

    return Promise.reject('Unknown method');
}

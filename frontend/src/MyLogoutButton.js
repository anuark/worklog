import React from 'react';
import { connect } from 'react-redux';
import { userLogout } from 'admin-on-rest';

const MyLogoutButton = ({ userLogout }) => (
    <button onClick={userLogout}>Logout</button>
);

export default connect(undefined, { userLogout })(MyLogoutButton);

import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import FlatButton from 'material-ui/FlatButton';
import RaisedButton from 'material-ui/RaisedButton';
import Dialog from 'material-ui/Dialog';
import {showNotification as showNotificationAction} from 'admin-on-rest';
import {push as pushAction} from 'react-router-redux';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';

class GenerateButton extends Component {
    state = {
        open: false,
    };

    handleOpen = () => {
        this.setState({open: true});
    };

    handleClose = () => {
        this.setState({open: false});
    };

    render() {
        const actions = [
            <FlatButton
                label="Cancel"
                primary={true}
                onClick={this.handleClose}
            />,
            <FlatButton
                label="Submit"
                primary={true}
                keyboardFocused={true}
                onClick={this.handleClose}
            />,
        ];
    
        return (
            <span>
            <FlatButton primary label="Generate PDF" onClick={this.handleOpen} icon={<FileCloudDownload />} />

            <Dialog
              title="Dialog With Actions"
              actions={actions}
              modal={false}
              open={this.state.open}
              onRequestClose={this.handleClose}
            >
              The actions in this window were passed in as an array of React objects.
            </Dialog>
            </span>
        );

        // return <FlatButton primary label="Generate PDF" onClick={console.log("clicked")} icon={<FileCloudDownload />} />;
    }
}
GenerateButton.propTypes = {
    push: PropTypes.func,
    record: PropTypes.object,
    showNotification: PropTypes.func
}

export default connect(null, {
    showNotification: showNotificationAction,
    push: pushAction
})(GenerateButton);

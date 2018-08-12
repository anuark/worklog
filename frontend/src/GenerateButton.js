import React, { Component } from 'react';
import PropTypes from 'prop-types';
import {connect} from 'react-redux';
import FlatButton from 'material-ui/FlatButton';
// import Dialog from 'material-ui/Dialog';
import {showNotification as showNotificationAction} from 'admin-on-rest';
import {push as pushAction} from 'react-router-redux';
import FileCloudDownload from 'material-ui/svg-icons/file/cloud-download';
// import DatePicker from 'material-ui/DatePicker'

class GenerateButton extends Component {
    state = {
        open: false,
        startDate: null,
        endDate: null
    };

    validate = (values) => {
        const errors = {};
        if (!values.startDate) {
            errors.startDate = ['Start date is required.'];
        }

        if (!values.endDate) {
            errors.endDate = ['End date is required.'];
        }

        return errors
    }

    handleOpen = () => {
        this.setState({open: true});
    };

    handleClose = () => {
        this.setState({open: false});
    };

    save = () => {
        console.log("submit");
        this.handleClose();
    }

    handleStartDateChange = (event, date) => {
        this.setState({startDate: date});
    }

    handleEndDateChange = (event, date) => {
        this.setState({endDate: date});
    }

    render() {
        // const actions = [
        //     <FlatButton
        //         label="Cancel"
        //         primary={true}
        //         onClick={this.handleClose}
        //     />,
        //     <FlatButton
        //         label="Generate"
        //         primary={true}
        //         keyboardFocused={true}
        //         onClick={this.handleClose}
        //     />,
        // ];
    
        // return (
        //     <span>
        //         <FlatButton primary label="Generate PDF" onClick={this.handleOpen} icon={<FileCloudDownload />} />
        //         <Dialog
        //         title="Generate pdf"
        //         actions={actions}
        //         modal={false}
        //         open={this.state.open}
        //         onRequestClose={this.handleClose}
        //         autoOk={true}
        //         >
        //             <SimpleForm resource="/" save={this.save} validate={this.validate}>
        //                 <TextInput source="name" />
        //                 <DateInput source="startDate" options={{
        //                     mode: 'landscape',
        //                     hintText: 'Start Date',
        //                     DateTimeFormat
        //                 }}/>
        //                 <DateInput label="End Date" source="endDate" options={{
        //                     mode: 'landscape',
        //                     hintText: 'Start Date',
        //                     DateTimeFormat
        //                 }}/>
        //                 {/* <DatePicker hintText="Start Date" /> */}
        //             </SimpleForm>
        //         </Dialog>
        //     </span>
        // );

        return <FlatButton primary href="/generate" label="Generate PDF" onClick={console.log("clicked")} icon={<FileCloudDownload />} />;
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

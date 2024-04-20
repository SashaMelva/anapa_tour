import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc, clearToastAc } from "../../redux/toastReducer";
import "./ToastComp.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class ToastComp extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
    this.handleClose = this.handleClose.bind(this)
  }

  handleClose() {
    this.props.clearToastAc()
  }

  
  render() {
      return <Toast className="Primary toast bg-warning text-dark" show={this.props.toastObj ? true : false} delay={this.props.toastObj?.delay ? this.props.toastObj?.delay : 3000} onClose={this.handleClose}>
        <Toast.Header className="text-body">Внимание!</Toast.Header>
        <Toast.Body className="text-body">{this.props.toastObj?.body}</Toast.Body>
      </Toast> 
  }
}

export default connect(mapStateToProps, {setToastAc, clearToastAc})(ToastComp)
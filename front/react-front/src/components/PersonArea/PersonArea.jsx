import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import "./PersonArea.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class PersonArea extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
  }

  
  render() {
      return <div>
        Личный кабинет
      </div>
  }
}

export default connect(mapStateToProps, {setToastAc})(PersonArea)
import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class MyMap extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
  }

  
  render() {
      return <div>
        карта
      </div>
  }
}

export default connect(mapStateToProps)(MyMap)
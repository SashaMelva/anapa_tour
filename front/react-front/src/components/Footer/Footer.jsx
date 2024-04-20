import React from "react";
import { connect } from 'react-redux';
import './Footer.css'

const mapStateToProps = (state) => {
  
}

const Footer = (props) => {
  return <div className="footer">
    Подвал
  </div>
}

export default connect(mapStateToProps)(Footer)
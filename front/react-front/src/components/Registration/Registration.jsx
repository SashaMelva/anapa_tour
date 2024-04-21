import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { setToastAc } from "../../redux/toastReducer";
import { Navigate } from 'react-router'
import "./Registration.css"
import { authIn } from "../../redux/registrationReducer";


const mapStateToProps = (state) => {
  return {
    isAuth: state.registration.isAuth
  }
}

class Registration extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      login: '',
      password: '',
      passwordAgain: '',
      isRegistration: false,
    };

    this.handleChange = this.handleChange.bind(this);
    this.handleSubmit = this.handleSubmit.bind(this);
  }

  handleChange(event, field) {
    this.setState({[field]: event.target.value});
  }

  handleSubmit(event) {
    event.preventDefault();
    const data = {
      login: this.state.login,
      password: this.state.password
    }

    if(!this.state.isRegistration) {
      authApi.login(data)
    } else {
      if(this.state.password === this.state.passwordAgain) {
        authApi.reg(data)
        this.props.setToastAc({body: "Регистрация прошла успешно!", delay: 2000})
        this.props.authIn()
      } else {
        this.props.setToastAc({body: "Пароли должны совпадать!", delay: 2000})
      }
    }
  }

  render() {
    const isReg = this.state.isRegistration
    if(this.props.isAuth) {
      return <Navigate to="/home"/>
    }
    return <div className={'authorization'}>
      <form onSubmit={this.handleSubmit} className="Primary">
        <label>
          Логин:
          <Form.Control 
            required 
            type="text" 
            value={this.state.login} 
            onChange={event => this.handleChange(event, 'login')} 
          />
        </label>
        <label>
          Пароль:
          <Form.Control 
            required 
            type="text" 
            value={this.state.password} 
            onChange={event => this.handleChange(event, 'password')} 
          />
        </label>
        {isReg && <label>
          Повторный пароль:
          <Form.Control 
            required 
            type="text" 
            value={this.state.passwordAgain} 
            onChange={event => this.handleChange(event, 'passwordAgain')} 
          />
        </label>}
        <Button variant="primary" type="submit" value="Отправить" style={{marginTop: "10px"}}>Войти</Button>
        <div style={{cursor: "pointer"}} onClick={() => this.setState({isRegistration: true})}>Регистрация</div>
      </form>
    </div>
  }
}

export default connect(mapStateToProps, { setToastAc, authIn })(Registration)
import React from "react";
import { connect } from 'react-redux';
import { Navbar, NavDropdown, Nav, Container, Form, Button } from "react-bootstrap";
import "./Header.css"
import { NavLink } from "react-router-dom";

const mapStateToProps = (state) => {
  
}

const Header = (props) => {
  return <Navbar expand="lg" className="bg-body-tertiary">
  <Container fluid>
    <Navbar.Brand href="#">Анапа(360)</Navbar.Brand>
    <Navbar.Toggle aria-controls="navbarScroll" />
    <Navbar.Collapse id="navbarScroll">
      <Nav
        className="me-auto my-2 my-lg-0 Nav"
        style={{ maxHeight: '100px' }}
        navbarScroll
      >
        <NavLink className={'text-body'} to="/home">Главная</NavLink>
        <NavLink className={'text-body'} to="/hotels">Бронь отелей</NavLink>
        <NavLink className={'text-body'} to="/restaurant">Рестораны(кафе)</NavLink>
        <NavLink className={'text-body'} to="/places">Места и события</NavLink>
        <NavDropdown title="Личный кабинет" id="nav-dropdown">
          <NavLink to='/personArea/myMap'>Моя карта</NavLink>
          <NavLink to='/personArea/people'>Люди</NavLink>
          <NavLink to='/personArea/favorites'>Избранное</NavLink>
          <NavLink to='/personArea/recomPlaces'>Рекомендуемые места</NavLink>
          <NavLink to='/personArea/treasureMap'>Карта сокровищ</NavLink>
      </NavDropdown>
      </Nav>
      <div>Сашка Супер Кодер</div>
    </Navbar.Collapse>
  </Container>
</Navbar>
}

export default connect(mapStateToProps)(Header)
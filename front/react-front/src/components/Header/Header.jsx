import React from "react";
import { connect } from 'react-redux';
import { Navbar, NavDropdown, Nav, Container, Form, Button } from "react-bootstrap";
import { NavLink } from "react-router-dom";
import "./Header.css"

const mapStateToProps = (state) => {
  return {

  }
}

const Header = (props) => {
  return <Navbar expand="lg" className="bg-body-tertiary">
  <Container fluid>
    <Navbar.Toggle aria-controls="navbarScroll" />
    <Navbar.Collapse id="navbarScroll">
      <Nav
        className="me-auto my-2 my-lg-0 Nav"
        style={{ maxHeight: '100px' }}
        navbarScroll
      >
        <NavLink className={'text-body'} to="/home">Главная</NavLink>
        <NavLink className={'text-body'} to="/hotels">Бронь отелей</NavLink>
        <NavLink className={'text-body'} to="/restaurants">Рестораны(кафе)</NavLink>
        <NavLink className={'text-body'} to="/places">Места и события</NavLink>
        <NavDropdown title="Личный кабинет" id="nav-dropdown">
          <NavLink to='/personArea/myMap'>Моя карта</NavLink>
          <NavLink to='/personArea/people'>Люди</NavLink>
          <NavLink to='/personArea/favorites'>Избранное</NavLink>
          <NavLink to='/personArea/recomPlaces'>Рекомендуемые места</NavLink>
          <NavLink to='/personArea/treasureMap'>Карта сокровищ</NavLink>
      </NavDropdown>
      </Nav>
      <div>alexandrNik</div>
    </Navbar.Collapse>
  </Container>
</Navbar>
}

export default connect(mapStateToProps)(Header);
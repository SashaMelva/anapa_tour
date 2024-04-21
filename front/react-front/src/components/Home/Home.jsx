import React from "react";
import { connect } from 'react-redux';
import { Form, Button, Col } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import { YMaps, Map, Route } from '@pbe/react-yandex-maps';
import { useRef, useEffect, useState } from "react";
import "./Home.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

const MyForm = () => {
  const [formData, setFormData] = useState({
    start: '',
    end: '',
    countDays: '',
    countAdults: '',
    countChild: '',
    money: '',
    availability: 'male',
    byCar: false,
    byPublicTr: false,
    byFoot: false,
    hasHome: false
  });

  const handleChange = (e) => {
    const { name, value, type, checked } = e.target;
    setFormData({
      ...formData,
      [name]: type === 'checkbox' ? checked : value,
    });
  };

  const handleSubmit = (e) => {
    e.preventDefault();
  };

  const onSubmiFormtBtn = () => {
    
  }

  return (
    <Form onSubmit={handleSubmit} className="myFormMain">
      <div className="formGroupCss">
        <Form.Group controlId="start">
          <Form.Label>Откуда</Form.Label>
          <Form.Control type="text" name="start" value={formData.start} onChange={handleChange} />
        </Form.Group>
        <Form.Group controlId="end">
          <Form.Label>Куда</Form.Label>
          <Form.Control type="text" name="end" value={formData.end} onChange={handleChange} />
        </Form.Group>
      </div>
      <div className="formGroupCss">
        <Form.Group controlId="countChild">
          <Form.Label>Кол-во детей </Form.Label>
          <Form.Control type="number" name="countChild" value={formData.countChild} onChange={handleChange} />
        </Form.Group>
        <Form.Group controlId="countAdults">
          <Form.Label>Кол-во взрослых</Form.Label>
          <Form.Control type="number" name="countAdults" value={formData.countAdults} onChange={handleChange} />
        </Form.Group>
        <Form.Group controlId="money">
          <Form.Label>Бюджет</Form.Label>
          <Form.Control type="number" name="money" value={formData.money} onChange={handleChange} />
        </Form.Group>
      </div>
      <div className="formGroupCss">
        <Form.Group controlId="transportation">
          <Form.Label>Доступность и выбор:</Form.Label>
          <Col>
            <Form.Check
              type="radio"
              label="Автомобилем"
              name="transportation"
              value="car"
              checked={formData.transportation === 'car'}
              onChange={handleChange}
            />
            <Form.Check
              type="radio"
              label="Автобусом"
              name="transportation"
              value="bus"
              checked={formData.transportation === 'bus'}
              onChange={handleChange}
            />
            <Form.Check
              type="radio"
              label="Пешком"
              name="transportation"
              value="walk"
              checked={formData.transportation === 'walk'}
              onChange={handleChange}
            />
          </Col>
        </Form.Group>
        <div style={{display: "flex", flexDirection: "column", gap: "10px"}}>
          <Form.Group controlId="startDate">
            <Form.Label>Заезд</Form.Label>
            <Form.Control type="number" name="age" value={formData.startDate} onChange={handleChange} />
          </Form.Group>
          <Form.Group controlId="startDate">
            <Form.Label>Выезд</Form.Label>
            <Form.Control type="number" name="age" value={formData.endDate} onChange={handleChange} />
          </Form.Group>
        </div>
      </div>
      <Form.Group controlId="hasHome">
        <Form.Check type="checkbox" name="hasHome" label="В жилье нет необходимости " checked={formData.hasHome} onChange={handleChange} />
      </Form.Group>
      <Button variant="primary" type="submit" onClick={onSubmiFormtBtn()}>
        Сформировать варианты отдыха
      </Button> 
    </Form>
  );
};

const Home = () => {
  const [start, setStart] = React.useState([]);
  const [end, setEnd] = React.useState([]);
  const [countDays, setCountDays] = React.useState([]);
  const [password, setPassword] = React.useState([]);

  let isRender = false
  useEffect(() => {
    const initializeMap = () => {
      // Создание карты
      const map = new window.ymaps.Map('map', {
        center: [44.900322, 37.381261], // Координаты центра карты
        zoom: 10, // Масштаб карты
      });

      // Создание маршрута
      window.ymaps.route([
        [44.844182, 37.377369], // Координаты начальной точки
        [44.916186, 37.393909], // Координаты конечной точки
      ]).then(function (route) {
        // Добавление маршрута на карту
        map.geoObjects.add(route);
      });
    };

    if (window.ymaps && !isRender) {
      isRender = true
      window.ymaps.ready(initializeMap);
    } else {
      console.error('API Яндекс.Карт не было загружено');
    }
  }, []);
  return <div className="main">
    <div className="main_left">
      <div className="main_leftHeader">
        <Button>Создать свой маршрут</Button>
        <Button>Готовые маршруты</Button>
      </div>
      <div id="map" style={{ width: '100%', height: '400px' }}></div>
      <div className="main_leftFooter">
        <Button>В городе </Button>
        <Button>По пути в город</Button>
      </div>
      </div>
      <div className="main_right">
        <MyForm/>
      </div>
    </div>
};

export default connect(mapStateToProps, {setToastAc})(Home);
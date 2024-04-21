import React from "react";
import { connect } from 'react-redux';
import { Form, Button, Col } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import { YMaps, Map, Route } from '@pbe/react-yandex-maps';
import { useRef, useEffect, useState } from "react";
import "./MainRouteOptions.css"
import { Navigate } from "react-router-dom";

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

const CameraCards = () => {

  return (
    <div className="camerCards_content">
      <div className="cameraCards_items">
        <div className="cameraCards_item">
          <div className="cameraCards_itemImg"/>
          <div className="cameraCards_itemBody">
            <div className="cameraCards_itemHeader">Золотая бухта</div>
            <div className="cameraCards_itemBusy">Процент загруженности 5%</div>
          </div>
        </div>
        <div className="cameraCards_item">
          <div className="cameraCards_itemImg"/>
          <div className="cameraCards_itemBody">
            <div className="cameraCards_itemHeader">Золотая бухта</div>
            <div className="cameraCards_itemBusy">Процент загруженности 5%</div>
          </div>
        </div>
        <div className="cameraCards_item">
          <div className="cameraCards_itemImg"/>
          <div className="cameraCards_itemBody">
            <div className="cameraCards_itemHeader">Золотая бухта</div>
            <div className="cameraCards_itemBusy">Процент загруженности 5%</div>
          </div>
        </div>
        <div className="cameraCards_item">
          <div className="cameraCards_itemImg"/>
          <div className="cameraCards_itemBody">
            <div className="cameraCards_itemHeader">Золотая бухта</div>
            <div className="cameraCards_itemBusy">Процент загруженности 5%</div>
          </div>
        </div>
        <div className="cameraCards_item">
          <div className="cameraCards_itemImg"/>
          <div className="cameraCards_itemBody">
            <div className="cameraCards_itemHeader">Золотая бухта</div>
            <div className="cameraCards_itemBusy">Процент загруженности 5%</div>
          </div>
        </div>
      </div>
    </div>
  );
};

const MyList = () => {
  const [toBusyProc, setToBusyProc] = React.useState(false);

  if(toBusyProc) {
    return <Navigate to="/busyProc"/>
  }

  return (
    <div className="routeCards_content">
      <h2>Варианты маршрута</h2>
      <div className="routeCards_items" style={{width: "380px", display: "flex", flexDirection: "column", gap: "10px"}}>
        <div className="routeCards_item" >
          <div className="routeCadrs_itemBlock">
            <div className="routeCards_itemText">
              Повседневная практика показывает, что поПовседневная практика Повседневная практика показывает, 
            </div>
            <div className="routeCards_itemFooter">
              <div className="routeCards_itemInfo">
                <p>минимальная бюджет</p>
                <div>
                  <p style={{borderRight: "1px solid black", paddingRight: "5px", marginRight: "5px"}}>4 дня</p>
                  <p>700 км</p>
                </div>
              </div>
              <div style={{display: "flex", alignItems: "center", border: "1px solid black", borderRadius: "10px", padding: "0 5px"}}>27 000 р.</div>
            </div>
          </div>
        </div>
        <div className="routeCards_item" >
            <div className="routeCadrs_itemBlock">
              <div className="routeCards_itemText">
                Повседневная практика показывает, что поПовседневная практика Повседневная практика показывает, 
              </div>
              <div className="routeCards_itemFooter">
                <div className="routeCards_itemInfo">
                  <p>минимальная бюджет</p>
                  <div>
                    <p style={{borderRight: "1px solid black", paddingRight: "5px", marginRight: "5px"}}>4 дня</p>
                    <p>700 км</p>
                  </div>
                </div>
                <div style={{display: "flex", alignItems: "center", border: "1px solid black", borderRadius: "10px", padding: "0 5px"}}>27 000 р.</div>
              </div>
            </div>
          </div>
          <div className="routeCards_item" >
            <div className="routeCadrs_itemBlock">
              <div className="routeCards_itemText">
                Повседневная практика показывает, что поПовседневная практика Повседневная практика показывает, 
              </div>
              <div className="routeCards_itemFooter">
                <div className="routeCards_itemInfo">
                  <p>минимальная бюджет</p>
                  <div>
                    <p style={{borderRight: "1px solid black", paddingRight: "5px", marginRight: "5px"}}>4 дня</p>
                    <p>700 км</p>
                  </div>
                </div>
                <div style={{display: "flex", alignItems: "center", border: "1px solid black", borderRadius: "10px", padding: "0 5px"}}>27 000 р.</div>
              </div>
            </div>
          </div>
          <div className="routeCards_item" >
            <div className="routeCadrs_itemBlock">
              <div className="routeCards_itemText">
                Повседневная практика показывает, что поПовседневная практика Повседневная практика показывает, 
              </div>
              <div className="routeCards_itemFooter">
                <div className="routeCards_itemInfo">
                  <p>минимальная бюджет</p>
                  <div>
                    <p style={{borderRight: "1px solid black", paddingRight: "5px", marginRight: "5px"}}>4 дня</p>
                    <p>700 км</p>
                  </div>
                </div>
                <div style={{display: "flex", alignItems: "center", border: "1px solid black", borderRadius: "10px", padding: "0 5px"}}>27 000 р.</div>
              </div>
            </div>
          </div>
          <div className="routeCards_item" >
            <div className="routeCadrs_itemBlock">
              <div className="routeCards_itemText">
                Повседневная практика показывает, что поПовседневная практика Повседневная практика показывает, 
              </div>
              <div className="routeCards_itemFooter">
                <div className="routeCards_itemInfo">
                  <p>минимальная бюджет</p>
                  <div>
                    <p style={{borderRight: "1px solid black", paddingRight: "5px", marginRight: "5px"}}>4 дня</p>
                    <p>700 км</p>
                  </div>
                </div>
                <div style={{display: "flex", alignItems: "center", border: "1px solid black", borderRadius: "10px", padding: "0 5px"}}>27 000 р.</div>
              </div>
            </div>
          </div>
      </div>
    </div>
  );
};

const MainRouteOptions = () => {
  const [start, setStart] = React.useState([]);
  const [end, setEnd] = React.useState([]);
  const [countDays, setCountDays] = React.useState([]);
  const [password, setPassword] = React.useState([]);
  const [toBusyProc, setToBusyProc] = React.useState(false);

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
  return <div className="mainRouter main">
    <div className="mainRouter_right main_right" style={{width: `${toBusyProc ? "100%" : "50%"}`}}>
      <Button className="mainBusyProc" 
      style={{backgroundColor: `${toBusyProc ? "#484AA1" : "white"}`, 
      border: `${toBusyProc ? "none" : "2px solid #484AA1"}`, 
      color:`${toBusyProc ? "white" : "inherit"}`, marginRight: "auto", width: "380px", marginBottom: "10px"}} onClick={() => setToBusyProc(!toBusyProc)}>Посмотреть % загруженности в этих местах</Button>
      {toBusyProc ? <CameraCards/> : <MyList/>}
    </div>
    <div className="mainRouter_left main_left" style={{display: `${toBusyProc ? "none" : "flex"}`, alignItems: "center", justifyContent: "center"}}> 
      <div id="map" style={{ width: '100%', height: '100%' }}></div>
    </div>
    </div>
};

export default connect(mapStateToProps, {setToastAc})(MainRouteOptions);
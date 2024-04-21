import React, { useEffect, useRef, useState } from "react"
import { Map, Placemark, SearchControl, GeolocationControl } from "@pbe/react-yandex-maps"
import { Button } from "react-bootstrap"
import { setToastAc } from "../../redux/toastReducer"
import { addChests } from "../../redux/chestReducer"
import { connect } from "react-redux"
import "./MyMap.css"
import { authApi } from "../../api/authApi"

const mapStateToProps = (state) => {
  return {
    chests: state.chestObj.chests
  }
}

const MyMap = React.memo((props) => {
  const [markers, setMarkers] = React.useState([]);
  const [apiLoaded, setApiLoaded] = useState(false);
  const [placemarkCoords, setPlacemarkCoords] = React.useState(null); // Координаты Placemark


  // // Обработчик события клика на карте
  // const handleMapClick = React.useCallback((e) => {
  //   const clickedCoordinates = e.get('coords');
  //   setMarkers([...markers, clickedCoordinates]); // Добавляем новую координату в состояние маркеров
  // }, [markers]);



  useEffect(() => {
    authApi.getChest().then(res => props.addChests(res))
    const loadApi = async () => {
      await new Promise((resolve, reject) => {
        const script = document.createElement('script');
        script.src = 'https://api-maps.yandex.ru/2.1/?apikey=YOUR_API_KEY&lang=ru_RU';
        script.onload = resolve;
        script.onerror = reject;
        document.body.appendChild(script);
      });
      setApiLoaded(true);
    };

    loadApi();
  }, []);

  const changeMapPos = (position) => {
    setPlacemarkCoords(position)
  }

  const openChest = (chest) => {
    const index = props.chests.findIndex(obj => obj.pin.coordinat === chest.pin.coordinat);

    props.setToastAc({body: "Сундук Открылся!"})
    const chestElem = document.querySelector(`.btn${chest.organization.name + index}`)
    chestElem.classList.add('isVisibleBtn')
  }

  function convertStringToArray(str) {
    // Удаляем фигурные скобки и разделяем строку по запятой
    const parts = str.replace(/[{}]/g, '').split(',');
    
    // Преобразуем каждый элемент в числовое значение
    const result = parts.map(parseFloat);
    
    return result;
}

  if (!apiLoaded ) {
    return <div>Загрузка API...</div>;
  }
  return (
    <div className="map_body">
      <div className="map_balanceBlock">
        <p>Мои Бонусы: </p>
        <p className="map_balanceText">360</p>
        <div className="map_balanceImg"/>
      </div>
      <div className="map_mainBlock">
        <div className="map">
          <Map style={{width: "400px", height: "400px"}}
            defaultState={{ center: convertStringToArray(props.chests[0].pin.coordinat), zoom: 9 }}
            state={{ center: placemarkCoords ? placemarkCoords : convertStringToArray(props.chests[0].pin.coordinat), zoom: 10 }}>
            {props.chests && props.chests.map((chest, index) => (
              <Placemark key={index} geometry={convertStringToArray(chest.pin.coordinat)}/>
            ))}
          </Map>
        </div>
        <div className="chestsList">
          {props.chests.map((chest, index) => (
            <div className="chestsList_chest" onClick={() => changeMapPos(convertStringToArray(chest.pin.coordinat))}>
            <div className="chest_block">
              <div className="chest_img"/>
              <div className="chest_info">
                <div className="chest_header">
                  <div className="chest_headerText">{chest.organization.name}</div>
                  <Button className={`btn btn-success`} onClick={() => openChest(chest)}>{"Открыть"}</Button>
                </div>
                <div style={{display: "none"}} className={`chest_diskount btn${chest.organization.name + index}`}>{chest.action.name}</div>
              </div>
            </div>
          </div>
          ))}
        </div>
      </div>
    </div >
  )
})
export default connect(mapStateToProps, {setToastAc, addChests})(MyMap);
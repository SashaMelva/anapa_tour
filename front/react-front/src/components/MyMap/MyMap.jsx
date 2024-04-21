import React, { useEffect, useRef, useState } from "react"
import { Map, Placemark, SearchControl, GeolocationControl } from "@pbe/react-yandex-maps"
import { Button } from "react-bootstrap"
import { setToastAc } from "../../redux/toastReducer"
import { connect } from "react-redux"
import "./MyMap.css"

const mapStateToProps = (state) => {
  return {
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

  const onPlacemarkClick = (position) => {
    const chest = props.chests.find(obj => obj.position === position);
  }

  useEffect(() => {
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
    props.setToastAc({body: "Сундук Открылся!"})
  }

  if (!apiLoaded) {
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
            defaultState={{ center: props.chests[0].position, zoom: 9 }}
            state={{ center: placemarkCoords ? placemarkCoords : props.chests[0].position, zoom: 10 }}>
            {props.chests && props.chests.map((chest, index) => (
              <Placemark key={index} geometry={chest.position} onClick={() => onPlacemarkClick(chest.position)}/>
            ))}
          </Map>
        </div>
        <div className="chestsList">
          {props.chests.map((chest) => (
            <div className="chestsList_chest" onClick={() => changeMapPos(chest.position)}>
            <div className="chest_block">
              <div className="chest_img"/>
              <div className="chest_info">
                <div className="chest_header">
                  <div className="chest_headerText">Мацони</div>
                  <Button className="btn btn-success" onClick={() => openChest(chest)}>Открыть</Button>
                </div>
                <div className="chest_diskount">Скидка на чек 10%  при предъявлении QR-кода</div>
              </div>
            </div>
          </div>
          ))}
        </div>
      </div>
    </div >
  )
})
export default connect(mapStateToProps, {setToastAc})(MyMap);
import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import Card from "../Card/Card";
import "./Hotels.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class Hotels extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
  }


  render() {
      return <div className="hotels">
        <div className="hotels_items">
          <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
          <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
            <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
            <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
            <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
            <Card 
            name="Город мира" 
            distance="650 км. до моря" 
            bonuses="Wi-Fi - Бассейн - Парковка - Кондиционер"
            price="4 223р."
            priceText="Цена на 1 ночь"
            comments="20 отзывов"
            />
        </div>
      </div>
  }
}

export default connect(mapStateToProps, {setToastAc})(Hotels);
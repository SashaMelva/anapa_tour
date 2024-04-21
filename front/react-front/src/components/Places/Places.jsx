import React from "react";
import { connect } from 'react-redux';
import { Form, Button } from "react-bootstrap";
import { authApi } from "../../api/authApi";
import { Toast } from "react-bootstrap";
import { setToastAc } from "../../redux/toastReducer";
import Card from "../Card/Card";
import "./Places.css"

const mapStateToProps = (state) => {
  return {
    toastObj: state.toast.toastObj
  }
}

class Places extends React.Component {
  constructor(props) {
    super(props);
    this.state = {
      
    };
  }


  render() {
      return <div className="place">
        <div className="place_items">
          <Card 
            name="Фестиваль" 
            distance="08:00-22:00" 
            bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
            price="1 800р."
            priceText="Стоимость билета "
            />
          <Card 
            name="Фестиваль" 
            distance="08:00-22:00" 
            bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
            price="1 800р."
            priceText="Стоимость билета "
            />
          <Card 
            name="Фестиваль" 
            distance="08:00-22:00" 
            bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
            price="1 800р."
            priceText="Стоимость билета "
            />
          <Card 
            name="Фестиваль" 
            distance="08:00-22:00" 
            bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
            price="1 800р."
            priceText="Стоимость билета "
            />
          <Card 
          name="Фестиваль" 
          distance="08:00-22:00" 
          bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
          price="1 800р."
          priceText="Стоимость билета "
          />
          <Card 
          name="Фестиваль" 
          distance="08:00-22:00" 
          bonuses="ANNA ASTI, JONY, Big Baby Tape, Хаски, Баста, Леонид Агутин, FEDUK, ИВАНУШКИ, Jah Khalib и ещё 80 популярных артистов"
          price="1 800р."
          priceText="Стоимость билета "
          />

        </div>
      </div>
  }
}

export default connect(mapStateToProps, {setToastAc})(Places);
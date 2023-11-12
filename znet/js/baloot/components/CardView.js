import PropTypes from 'prop-types';
import React from "react";
import {
  SimpleGrid,
  Button,
  Card,
  Heading,
  CardHeader,
  CardBody,
  CardFooter,
  Text,
} from "@chakra-ui/react";
import { TrefleIcon, CarreauIcon, PiqueIcon, CoeurIcon } from "./Icons";

const CardIcon = ({couleur, size}) => {
  if (couleur == "Carreau") {
    return <CarreauIcon boxSize={size} color="red.500" />
  }
  if (couleur == "Pique") {
    return <PiqueIcon w={16} h={16} color="red.500" />
  }
  if (couleur == "Coeur") {
    return <CoeurIcon w={16} h={16} color="red.500" />
  }
  if (couleur == "Trefle") {
    return <TrefleIcon w={16} h={16} color="red.500" />
  }
}

function CardView({
  Genre,
  Couleur,
  onDragStart,
  onDragEnter,
  onDragEnd,
  onClick,
  width,
  height,
}) {
  return (
    <Card
      w={width}
      h={height}
      draggable
      onDragStart={onDragStart}
      onDragEnter={onDragEnter}
      onDragEnd={onDragEnd}
      onClick={onClick ? onClick.bind(this, Couleur, Genre) : null}
      size="sm"
    >
      <CardHeader>
        <Text textAlign={"left"} fontSize="22" flex="1">
          {Genre}
        </Text>
      </CardHeader>
      <CardBody align="center">
        <CardIcon couleur={Couleur} size={10} />
      </CardBody>
      <CardFooter>
        <Text textAlign={"right"} fontSize="22" flex="1">
          {Genre}
        </Text>
      </CardFooter>
    </Card>
  );
}

CardView.PropTypes = {
  Genre: PropTypes.string.isRequired,
  Couleur: PropTypes.string.isRequired,
  onDragStart: PropTypes.func,
  onDragEnter: PropTypes.func,
  onDragEnd: PropTypes.func,
  onClick: PropTypes.string.func,
  width: PropTypes.number,
  height: PropTypes.number,
}

CardView.defaultProps = {
  width: 128,
  height: 64
}

export default CardView;

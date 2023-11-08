import React from "react";
import { Box, Button, ButtonGroup } from "@chakra-ui/react";

function TakesGroupView({ onClickHandler, playerID, takes }) {
  console.log(takes)
  return (
    <Box m="2">
      <ButtonGroup size="sm" isAttached variant="outline" spacing={1}>
        {takes.map((t) => {
          return (
            <Button name={t} onClick={(e) => onClickHandler(t, playerID)}>
              {t}
            </Button>
          );
        })}
      </ButtonGroup>
    </Box>
  );
}

export default TakesGroupView;

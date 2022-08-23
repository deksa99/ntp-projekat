import React from "react";
import { Box, Tab } from "@mui/material";
import { TabList, TabPanel, TabContext } from "@mui/lab";
import CarServiceCatalogTable from "./CarServiceCatalogTable";

const CarServiceDetailsTabs = ({ cs }) => {
  const [value, setValue] = React.useState("1");

  const handleChange = (_, newValue) => {
    setValue(newValue);
  };

  console.log(cs);

  return (
    <Box sx={{ width: "100%", typography: "body1" }}>
      <TabContext value={value}>
        <Box sx={{ borderBottom: 1, borderColor: "divider" }}>
          <TabList onChange={handleChange} centered variant="fullWidth">
            <Tab label="Usluge" value="1" />
            <Tab label="Recenzije" value="2" />
          </TabList>
        </Box>
        <TabPanel value="1">
          <CarServiceCatalogTable cs={cs} />
        </TabPanel>
        <TabPanel value="2">Ocene</TabPanel>
      </TabContext>
    </Box>
  );
};

export default CarServiceDetailsTabs;

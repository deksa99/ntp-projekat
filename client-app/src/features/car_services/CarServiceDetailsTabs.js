import React from "react";
import { Box, Tab } from "@mui/material";
import { TabList, TabPanel, TabContext } from "@mui/lab";
import CarServiceCatalogTable from "./CarServiceCatalogTable";
import CarServiceReviewsTable from "./CarServiceReviewsTable";

const CarServiceDetailsTabs = ({ cs }) => {
  const [value, setValue] = React.useState("1");

  const handleChange = (_, newValue) => {
    setValue(newValue);
  };

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
        <TabPanel value="2">
          <CarServiceReviewsTable cs={cs} />
        </TabPanel>
      </TabContext>
    </Box>
  );
};

export default CarServiceDetailsTabs;

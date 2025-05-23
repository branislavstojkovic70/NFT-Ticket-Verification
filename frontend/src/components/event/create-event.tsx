import {
    Button,
    MenuItem,
    TextField,
    InputAdornment,
    IconButton,
  } from "@mui/material";
  import { useFormik } from "formik";
  import * as yup from "yup";
  import { useNavigate } from "react-router-dom";
  import { useState } from "react";
  import toast from "react-hot-toast";
  import api from "../../config/axios-config.tsx";
  import "../auth/login-registration-form.css";
  
  const eventTypes = [
    { label: "Music", value: "music" },
    { label: "Conference", value: "conference" },
  ];
  
  const validationSchema = yup.object({
    title: yup.string().required("Title is required"),
    location: yup.string().required("Location is required"),
    type: yup.string().oneOf(["music", "conference"]).required("Type is required"),
    dateStart: yup.date().required("Start date is required"),
    dateEnd: yup
      .date()
      .min(yup.ref("dateStart"), "End date must be after start date")
      .required("End date is required"),
    description: yup.string().required("Description is required"),
    tags: yup.string().required("Tags (as JSON array) are required"),
    numberOfTickets: yup
      .number()
      .min(1, "Must be at least 1 ticket")
      .required("Number of tickets is required"),
  });
  
  export default function CreateEventForm() {
    const navigate = useNavigate();
    const [organizerId] = useState(localStorage.getItem("organizer_id") || ""); // Pretpostavka
  
    const formik = useFormik({
      initialValues: {
        title: "",
        location: "",
        type: "",
        dateStart: "",
        dateEnd: "",
        description: "",
        tags: "",
        numberOfTickets: 100,
      },
      validationSchema: validationSchema,
      onSubmit: (values) => {
        let tagsJson;
        try {
          tagsJson = JSON.parse(values.tags);
          if (!Array.isArray(tagsJson)) throw new Error();
        } catch {
          toast.error("Tags must be a valid JSON array");
          return;
        }
  
        api
          .post("/events", {
            title: values.title,
            location: values.location,
            type: values.type,
            date_start: values.dateStart,
            date_end: values.dateEnd,
            description: values.description,
            tags: tagsJson,
            number_of_tickets: values.numberOfTickets,
            organizer_id: organizerId,
          })
          .then(() => {
            toast.success("Event created!");
            navigate("/events");
          })
          .catch((error) => {
            toast.error(error.response?.data?.error || "Error creating event");
          });
      },
    });
  
    return (
      <div id="form-container" style={{ width: "60%", padding: "5%" }}>
        <div className="headings">
          <h2>Create new event</h2>
        </div>
  
        <form onSubmit={formik.handleSubmit}>
          <TextField
            fullWidth
            id="title"
            name="title"
            label="Title"
            value={formik.values.title}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.title && Boolean(formik.errors.title)}
            helperText={formik.touched.title && formik.errors.title}
          />
  
          <TextField
            fullWidth
            id="location"
            name="location"
            label="Location"
            value={formik.values.location}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.location && Boolean(formik.errors.location)}
            helperText={formik.touched.location && formik.errors.location}
          />
  
          <TextField
            fullWidth
            select
            id="type"
            name="type"
            label="Event Type"
            value={formik.values.type}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.type && Boolean(formik.errors.type)}
            helperText={formik.touched.type && formik.errors.type}
          >
            {eventTypes.map((option) => (
              <MenuItem key={option.value} value={option.value}>
                {option.label}
              </MenuItem>
            ))}
          </TextField>
  
          <TextField
            fullWidth
            id="dateStart"
            name="dateStart"
            label="Start Date"
            type="datetime-local"
            InputLabelProps={{ shrink: true }}
            value={formik.values.dateStart}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.dateStart && Boolean(formik.errors.dateStart)}
            helperText={formik.touched.dateStart && formik.errors.dateStart}
          />
  
          <TextField
            fullWidth
            id="dateEnd"
            name="dateEnd"
            label="End Date"
            type="datetime-local"
            InputLabelProps={{ shrink: true }}
            value={formik.values.dateEnd}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.dateEnd && Boolean(formik.errors.dateEnd)}
            helperText={formik.touched.dateEnd && formik.errors.dateEnd}
          />
  
          <TextField
            fullWidth
            multiline
            id="description"
            name="description"
            label="Description"
            value={formik.values.description}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.description && Boolean(formik.errors.description)}
            helperText={formik.touched.description && formik.errors.description}
          />
  
          <TextField
            fullWidth
            id="tags"
            name="tags"
            label='Tags (e.g. ["tech", "blockchain"])'
            value={formik.values.tags}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.tags && Boolean(formik.errors.tags)}
            helperText={formik.touched.tags && formik.errors.tags}
          />
  
          <TextField
            fullWidth
            type="number"
            id="numberOfTickets"
            name="numberOfTickets"
            label="Number of Tickets"
            value={formik.values.numberOfTickets}
            onChange={formik.handleChange}
            onBlur={formik.handleBlur}
            error={formik.touched.numberOfTickets && Boolean(formik.errors.numberOfTickets)}
            helperText={formik.touched.numberOfTickets && formik.errors.numberOfTickets}
          />
  
          <Button
            sx={{ fontSize: "20px", textTransform: "capitalize", marginTop: "1rem" }}
            variant="contained"
            fullWidth
            type="submit"
          >
            Create Event
          </Button>
        </form>
      </div>
    );
  }
  